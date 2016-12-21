// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package fsm_test

import (
	"testing"

	"github.com/go-trellis/fsm"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	namespace  = "namespace1"
	namespace2 = "namespace2"
)

func TestFSM(t *testing.T) {
	Convey("testing get target transaction", t, func() {
		f := fsm.New()
		Convey("will return nil transaction", func() {
			Convey("when system not exists", func() {
				tran := fsm.New().GetTargetTranstion("namespace", "", "")
				So(tran, ShouldBeNil)
			})

			Convey("when added nil or empty namespace", func() {
				f.Add(nil)
				tran := f.GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldBeNil)
				f.Add(&fsm.Transaction{
					Namespace: "",
				})
				tran = f.GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldBeNil)
			})
			Convey("when added a transation without target", func() {
				f.Add(&fsm.Transaction{
					Namespace:     namespace,
					CurrentStatus: "status1",
					Event:         "event1",
					TargetStatus:  "",
				})
				tran := f.GetTargetTranstion("namespace2", "status1", "event1")
				So(tran, ShouldBeNil)
				tran = f.GetTargetTranstion(namespace, "status1", "event2")
				So(tran, ShouldBeNil)
			})
			Convey("when added a transation and input namespace not exists", func() {
				f.Add(&fsm.Transaction{
					Namespace:     namespace,
					CurrentStatus: "status1",
					Event:         "event1",
					TargetStatus:  "status2",
				})
				tran := f.GetTargetTranstion("namespace2", "status1", "event1")
				So(tran, ShouldBeNil)
				tran = f.GetTargetTranstion(namespace, "status1", "event2")
				So(tran, ShouldBeNil)
			})

			Convey("when remove a transation by information", func() {
				f.RemoveByTransaction(&fsm.Transaction{
					Namespace:     namespace,
					CurrentStatus: "status1",
					Event:         "event1",
				})
				tran := f.GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldBeNil)
			})
		})
		Convey("will return tran", func() {

			Convey("when tran exists", func() {
				f.Add(&fsm.Transaction{
					Namespace:     namespace,
					CurrentStatus: "status1",
					Event:         "event1",
					TargetStatus:  "status2",
				})
				f.Add(&fsm.Transaction{
					Namespace:     namespace2,
					CurrentStatus: "status1",
					Event:         "event1",
					TargetStatus:  "status3",
				})
				tran := fsm.New().GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldNotBeNil)
				So(tran.TargetStatus, ShouldEqual, "status2")
				tran = fsm.New().GetTargetTranstion(namespace2, "status1", "event1")
				So(tran, ShouldNotBeNil)
				So(tran.TargetStatus, ShouldEqual, "status3")
			})

			Convey("when remove invalid transation", func() {
				f.RemoveByTransaction(&fsm.Transaction{
					Namespace:     namespace,
					CurrentStatus: "",
					Event:         "event1",
				})
				tran := f.GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldNotBeNil)
				So(tran.TargetStatus, ShouldEqual, "status2")
			})
		})
		Convey("remove namespace", func() {
			Convey("remove empty namespace and will get transaction", func() {
				f.RemoveNamespace("")
				tran := fsm.New().GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldNotBeNil)
				So(tran.TargetStatus, ShouldEqual, "status2")
				tran = fsm.New().GetTargetTranstion(namespace2, "status1", "event1")
				So(tran, ShouldNotBeNil)
				So(tran.TargetStatus, ShouldEqual, "status3")
			})

			Convey("remove namespace and will get namespace2's transaction", func() {
				f.RemoveNamespace(namespace)
				tran := fsm.New().GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldBeNil)
				tran = fsm.New().GetTargetTranstion(namespace2, "status1", "event1")
				So(tran, ShouldNotBeNil)
				So(tran.TargetStatus, ShouldEqual, "status3")
			})
		})
		Convey("remove all transaction", func() {
			f.Remove()
			Convey("when get nil", func() {
				tran := fsm.New().GetTargetTranstion(namespace, "status1", "event1")
				So(tran, ShouldBeNil)
				tran = fsm.New().GetTargetTranstion(namespace2, "status1", "event1")
				So(tran, ShouldBeNil)
			})
		})
	})
}
