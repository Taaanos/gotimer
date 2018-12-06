# gotimer

## Overview

A simple way to time functions.

## Installation 

`go get -u github.com/Taaanos/gotimer`

## Usage

Use on the top of function you want to time with `defer`.

e.g 

`defer gotimer.TimeTrack(time.Now(),"function name")`

[source](https://blog.stathat.com/2012/10/10/time_any_function_in_go.html)
