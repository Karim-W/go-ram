package main

import "math/rand"

type Insults []string

func (i Insults) GetRandomInsult() string {
	return i[rand.Intn(len(i))]
}

var InsultsList = Insults{
	`My gran could do better! And she is dead!`,
	`For what we are about to eat, may the Lord make us truly not vomit.`,
	`You are getting your knickers in a twist! Calm down!`,
	`This lamb is so undercooked, it is following Mary to school!`,
	`This pizza is so disgusting, if you take it to Italy you will get arrested.`,
	`There is enough garlic in here to kill every vampire in Europe.`,
	`Why did the chicken cross the road? Because you did not fucking cook it!`,
	`You put so much ginger in this, it is a Weasley.`,
	`Do not just stand there like a big fucking muffin`,
	`Stop taking things personally.`,
	`I wouldn't trust you running a bath let alone a restaurant.`,
	`This fish is so raw, he's still finding Nemo.`,
	`You added so much salt and pepper I can hear the dish singing Push It.`,
	`Hey, panini head, are you even listening to me?`,
	`I'll get you more pumpkin, I'll ram it right up your ass`,
	`Get your shit together!`,
	`I swim like a fish, and I have an amazing kick`,
	`You used so much oil, the U.S. want to invade the fucking plate.`,
	`You've got to kiss arse to get somewhere, to learn. Clock-watchers are no good at kissing arse.`,
	`They say cats have nine lives. I've had 12 already and I don't know how many more I'll have.`,
	`This crab is so undercooked I can still hear it singing Under the Sea!'`,
	`This soufflé has sunk so badly James Cameron wants to make a film about it.`,
	`I've never, ever, ever, ever, ever met someone I believe in as little as you.`,
	`Why did the chicken cross the road? Because you didn't fucking cook it!`,
	`What are you? An idiot sandwich.`,
	`This squid is so undercooked I can still hear it telling Spongebob to fuck off.`,
	`This chicken is so uncooked that a skilled vet could still save him!`,
	`This fish is so Frozen that it is still singing "Let it Go!"`,
	`Just, just just - what do you want, a fucking medal?`,
	`and you, open those big eyes and watch what the fuck this guy is doing!!!`,
	`Oh come on. You give me them anaemic bits of shit, I'll fucking throw them up your arse sideways. [kicks a bin] Where's your fucking brain? I just cannot believe this! Can we have the two main courses TOGETHER?!!! [kicks the bins] SHIT!!! (groans)`,
	`Hey. Hey, you. Hey. Come here. Let me whisper something very important in your ear. Very important: FUCK OFF.`,
	`I wish you’d jump in the oven. That would make my life a whole lot easier.`,
	`Salmonella is not an ingridient.`,
}

// mix of gordon ramasy insults and inspired Insults
var EnduserInspiredInsults = Insults{
	"What are you? an idiot end-user",
	`I wouldn't trust you running a bath let alone using my app.`,
	`Hey, panini head, are you even listening to me?`,
	`Get your shit together!`,
	`I've never, ever, ever, ever, ever met someone I believe in as little as you.`,
	`I wish you’d jump in the oven. That would make my life a whole lot easier.`,
	`Just, just just - what do you want, a fucking medal?`,
	`Fuck off`,
}
