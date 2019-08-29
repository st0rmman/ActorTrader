package main

import (
	"./Const"
	"./Quote"
	"./gosiris"
)

func main() {
	//Init a local actor system
	gosiris.InitActorSystem(gosiris.SystemOptions{
		ActorSystemName: "ActorSystem",
	})

	//Create an actor
	quoteManger := Quote.NewQuoteManger()
	//Close an actor
	defer quoteManger.Base.Close()

	//Create an actor
	childActor := gosiris.Actor{}
	//Close an actor
	defer childActor.Close()
	//Register a reaction to event types ("message" in this case)
	childActor.React("message", func(context gosiris.Context) {
		context.Self.LogInfo(context, "Received %v\n", context.Data)
	})

	//Register an actor by spawning it
	gosiris.ActorSystem().SpawnActor(&quoteManger.Base, "childActor", &childActor, nil)

	//Retrieve actor references
	parentActorRef, _ := gosiris.ActorSystem().ActorOf(Const.QuoteManger)
	childActorRef, _ := gosiris.ActorSystem().ActorOf("childActor")

	//Send a message from one actor to another (from parentActor to childActor)
	childActorRef.Tell(gosiris.EmptyContext, "message", "Hi! How are you?", parentActorRef)
}