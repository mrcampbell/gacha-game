---
title: "Overkill: The Game, An Intro"
description: Let's make a game with a ton of tech
updatedAt: 2021-01-08
---

# Overkill: The Game, An Intro

I learn best when I'm building something I love, and from the start I've loved making games.

It all started with a Pokemon Battle clone in C++.  I didn't know about custom classes, so I used an Array to store each pokemon's values.  I'd like to hope I've come a ways since then, but that spark of making games has remained just as strong as it started.

In this series of posts, we're going to be making a Gacha Game.  We're going to start with a Monolith in `Go` and `Postgres`, and then break up that monolith into microservices that use Rust, Elixir, Node.js, and more!  Why?  For the sake of learning.  Will it be Overkill? Yep, but that's the name of the game.

First off, what is a Gacha Game?

```
Gacha games are video games that implement the gacha (capsule-toy vending machine) mechanic. This is somewhat similar to loot boxes,
inducing players to spend in-game currency to receive a random virtual item. Most of these games are free-to-play mobile games,
where the gacha serves as an incentive to spend real-world money.
```
source: [Wikipedia](https://en.wikipedia.org/wiki/Gacha_game)

That sucks.  The game creators goal is to provide a free game, keep people in as long as possible, force them into bottlenecks to squeeze money out of them, and they're not meant to be won. Once you've won, you're not incentivized to pay money, so they keep you on an endless journey.  There are "loot boxes" where you can gamble your money in an attempt to unlock characters and upgrades, but the chances of jackpots are not guaranteed and slimmer than advertized.  

Think about it - how many people do you know who have won Candy Crush?  How much has someone spent who has played it longer than a week?

I've been playing `Star Wars: Heroes`, and as much as I love it, I see the developers trying to squeeze dollars out of me, there's no hope of ever "beating it", and that needs to be fixed.  It should be said that I have not spent ONE dollar playing the game, and never will.

The idea of Gacha Games are fun, but the execution of them is corrupt.

Let's fix that!

## The Project


#### The Basic Game Mechanics

We're going to build a 6v6 Battle Game that uses a lot of traditional game mechanics, such as Classes (Healer, Tank, Attacker, Support), Elemental Advantage (Fire, Grass, Water), and then slowly accumulate features like daily objectives, progression to levels of increasing difficulty, leveling up and upgrading fighters, we'll create tooling to add content from an administrative dashboard, and more!

#### Improvement on Existing Gacha Games

We'll try to solve some of the less savory aspects of gacha games, like using alternative forms of income (background ads, and watching ads rather than purchasing), and mechanics we can use in place of gambling.  Rather than leave it to chance, I want to make it so skill is involved in unlocking characters, not a roll of the dice.  We'll create a "Challenge System" that you can plug whatever you want into it - complete a battle, solve a math problem, select a multiple choice flash card, you name it.  The quicker they solve the "challenge" the more lucrative their reward is.

#### Software 

Technically, we'll add monitoring, tracing, logging, and examine architectural design, discussing the pros and cons of a few different approaches.

We'll be exploring Go, Rust, Elixir, C#, and Node.js (and possibly more), microservices, SQL, NoSQL, React, Kubernetes, Docker, Prometheus, Grafana, gRPC, GraphQL, REST APIs, caching with Redis, and deploying our code with Terraform, and much, much more!\

Some of the tutorials will be "Main Quest", where it's using the standard or more basic approach to a problem, but there will be "Side Quests" along the way that use an alternate technology that can be a drop-in replacement for the "Main Quest" solution.  An example of this is, rather than using Websockets, we can use GraphQL Subscriptions, gRPC/WebRTC Streams, Long-Polling, etc.  We'll store these in a Sidequest `mod`/`app`/`package`/`directory`, so if you hop in during the middle of a "starter repo", you may see those on the side.  None of them need to be done to complete the game, so pick and choose as you like.

But this is a "choose your own adventure".  If some of those don't interest you, feel free to drop in at any stage, pick up the "starter repo" for your chapter, and go from there!

#### Business Administration

Business-wise, we'll be discussing how to measure and respond to user metrics, discussing statics like churn rate, customer lifetime value.  We'll install analytics tools to observe user activity, and use that information to make informed decisions about how we'll tweak the game.

We'll do A/B testing to perform experiments and see which updates to the UI lead to increased activity, making incremental improvements to the user experience.

## The Steps We'll Take

This tutorial will be broken up into steps, each with a starter repository.  If you want to skip a step, or are only interested in certain sections, each will have a repo you can clone from the previous step.

We're going to build the following in the following order:

```
0. Design the Game
  - A. I'll explain all the mechanics of the Proof of Concept so we 
      have a clear picture of what this game will look and feel like

1. A Go Monolith that uses an "inmemory database" as our proof of concept Backend API
  - A. Create a Fighter Service that serves up our game's Characters
  - B. Build out the Battle Mechanics
  - C. Write basic tests for our API
  - D. Serve our resources over HTTP with the Gin Framework

2. Begin on our Frontend that will use React and Chakra UI
  - A. Create a "View Fighters" page that calls our backend and displays
      the Fighter Details
  - B. Build a simple Battle Page stub (that displays details, but 
      doesn't function)

3. Build out the Battle System
  - A. In the Backend, we'll implement a Websocket service that we'll
      use to fight battles
  - B. Then, we'll use that Websocket in the Frontend to fight those    
      battles.


At this point, we'll consider our "Proof of Concept" complete, and move
on to splitting it into Microservices and connecting to databases and
caching systems!

To be continued...
```
