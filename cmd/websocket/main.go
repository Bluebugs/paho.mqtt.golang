/*
 * Copyright (c) 2021 Fynelabs and others.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * and Eclipse Distribution License v1.0 which accompany this distribution.
 *
 * The Eclipse Public License is available at
 *    https://www.eclipse.org/legal/epl-2.0/
 * and the Eclipse Distribution License is available at
 *   http://www.eclipse.org/org/documents/edl-v10.php.
 *
 * Contributors:
 *    Cedric Bail
 */

package main

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	//	wait := make(chan struct{})

	opts := mqtt.NewClientOptions()
	opts.AddBroker("ws://test.mosquitto.org:8080/")
	opts.SetClientID("gopherjs.deadlock")
	opts.AutoReconnect = true

	client := mqtt.NewClient(opts)

	//	go func() {
	<-client.Connect().Done()

	token := client.Subscribe("gopherjs/tests/attributes", 1, func(client mqtt.Client, msg mqtt.Message) {
		client.Unsubscribe("gopherjs/tests/attributes")
		//			wait <- struct{}{}
	})
	<-token.Done()

	token = client.Publish("gopherjs/tests/attributes", 1, false, "true")
	<-token.Done()
	//	}()

	//	<-wait
}
