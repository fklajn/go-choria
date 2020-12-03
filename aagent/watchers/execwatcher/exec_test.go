package execwatcher

import (
	"encoding/json"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AAgent/Watchers/ExecWatcher")
}

var _ = Describe("ExecWatcher", func() {
	Describe("setProperties", func() {
		It("Should parse valid properties", func() {
			w := &Watcher{}

			prop := map[string]interface{}{
				"command":                   "cmd",
				"timeout":                   "1.5s",
				"environment":               []string{"key1=val1", "key2=val2"},
				"suppress_success_announce": "true",
			}
			Expect(w.setProperties(prop)).ToNot(HaveOccurred())
			Expect(w.command).To(Equal("cmd"))
			Expect(w.timeout).To(Equal(1500 * time.Millisecond))
			Expect(w.environment).To(Equal([]string{"key1=val1", "key2=val2"}))
			Expect(w.suppressSuccessAnnounce).To(BeTrue())
		})

		It("Should handle errors", func() {
			w := &Watcher{}
			err := w.setProperties(map[string]interface{}{})
			Expect(err).To(MatchError("command is required"))
		})

		It("Should enforce 1 second intervals", func() {
			w := &Watcher{}
			err := w.setProperties(map[string]interface{}{
				"command": "cmd",
				"timeout": "0",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(w.command).To(Equal("cmd"))
			Expect(w.timeout).To(Equal(time.Second))
		})
	})

	Describe("CurrentState", func() {
		var (
			mockctl     *gomock.Controller
			mockMachine *MockMachine
			watcher     *Watcher
			now         time.Time
		)

		BeforeEach(func() {
			mockctl = gomock.NewController(GinkgoT())
			mockMachine = NewMockMachine(mockctl)

			now = time.Unix(1606924953, 0)
			mockMachine.EXPECT().Name().Return("exec").AnyTimes()
			mockMachine.EXPECT().Identity().Return("ginkgo").AnyTimes()
			mockMachine.EXPECT().InstanceID().Return("1234567890").AnyTimes()
			mockMachine.EXPECT().Version().Return("1.0.0").AnyTimes()
			mockMachine.EXPECT().TimeStampSeconds().Return(now.Unix()).AnyTimes()

			watcher = &Watcher{command: "/bin/sh", previous: Success, previousRunTime: time.Second, machine: mockMachine, name: "ginkgo"}
		})

		AfterEach(func() {
			mockctl.Finish()
		})

		It("Should be a valid state", func() {
			cs := watcher.CurrentState()
			csj, err := cs.(*StateNotification).JSON()
			Expect(err).ToNot(HaveOccurred())

			event := map[string]interface{}{}
			err = json.Unmarshal(csj, &event)
			Expect(err).ToNot(HaveOccurred())
			delete(event, "id")

			Expect(event).To(Equal(map[string]interface{}{
				"time":        "2020-12-02T16:02:33Z",
				"type":        "io.choria.machine.watcher.exec.v1.state",
				"subject":     "ginkgo",
				"specversion": "1.0",
				"source":      "io.choria.machine",
				"data": map[string]interface{}{
					"command":           "/bin/sh",
					"previous_outcome":  "success",
					"previous_run_time": float64(time.Second),
					"id":                "1234567890",
					"identity":          "ginkgo",
					"machine":           "exec",
					"name":              "ginkgo",
					"protocol":          "io.choria.machine.watcher.exec.v1.state",
					"type":              "exec",
					"version":           "1.0.0",
					"timestamp":         float64(now.Unix()),
				},
			}))
		})
	})
})
