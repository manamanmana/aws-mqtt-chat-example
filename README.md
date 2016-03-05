# aws-mqtt-chat-example

## What is this

This is just a small exmaple chat application using AWS IoT MQTT Broker.

## Preparation

### Make "things" on AWS IoT

1. From AWS Console open [AWS IoT](https://ap-northeast-1.console.aws.amazon.com/iot/home?region=ap-northeast-1#/dashboard?editor=thing)
![aws things console](https://raw.githubusercontent.com/wiki/manamanmana/aws-mqtt-chat-example/images/aws-things.PNG)

2. Put the "things" name into the text field of "Name" and click "Create".
![aws things console2](https://raw.githubusercontent.com/wiki/manamanmana/aws-mqtt-chat-example/images/aws-things2.PNG)
![aws things result](https://raw.githubusercontent.com/wiki/manamanmana/aws-mqtt-chat-example/images/aws-thingsi-result.PNG)

3. Select the thing which you created in the bottom, and click "Connect a device" button in the right pane.

4. On the next page, select "NodeJS" and click "Generate certificate and policy".

5. A few seconds later you can see "Download public key", "Download private key" and "Download certificate" links in the page. Download the certs from the 3 links.

6. Click "Confirm & start connecting". You can see the "Sample code" JSON in the page. Copy it and save it as a file (ex: ExampleChat.json) This file is used for the Chat Clients later.







