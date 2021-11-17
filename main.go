package main

import (
	"fmt"
	"io/ioutil"
	"log"

	examplepb "github.com/angelAL21/proto/example"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

func main() {
	sm := doSimple()

	rnwDemo(sm)
	jsonDemo(sm)

}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.EnumDayOfTheWeek_MONDAY,
	}
	em.DayOfTheWeek = enumpb.DayOfTheWeek_TUESDAY
	fmt.Println(em)
}

//jsonDemo runs the code of json and prints json
func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &examplepb.MyMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("sucessfully created proto struct: ", sm2)
}

//convert out string into json
func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("cant convert to JSON", err)
		return ""
	}
	return out
}

//unmarshal the string into pb
func fromJSON(in string, pb proto.Message) string {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("couldnt unmarshal the JSON into pb struct", err)
	}
	return in
}

func rnwDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &examplepb.MyMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

//writeToFile writes our doSimple func into a bin archive
func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("can serialise to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("can write to file", err)
		return err
	}
	fmt.Println("DATA HAS BEEN WRITTEN")
	return nil
}

//readFromFile reads from the simple.bin and prints it
func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("something went wron while reading", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("could not put the bytes into pb struct", err)
		return err
	}
	return nil
}

func doSimple() *examplepb.MyMessage {
	sm := examplepb.MyMessage{ //examplepb helps to know that comes from to pb file
		Id:          1243,
		FirstName:   "John",
		IsValidated: true,
	}
	fmt.Println(sm)

	sm.FirstName = "angel"

	fmt.Println(sm)
	fmt.Println("the id is: ", sm.GetId())

	return &sm

}
