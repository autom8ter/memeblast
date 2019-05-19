# MEMEBLAST

Send sms blasts of memes from a randomized set of phone numbers

**FOR EXPERIMENTATION PURPOSES ONLY**

    go get github.com/autom8ter/memeblast

## Requirements
- a twilio account with active numbers
- target phone numbers
- a config file named memeblast.yaml in $PWD 

## Example Config ($PWD/memeblast.yaml)

```text
## number of texts to send to each number(to) from a randomized number
count: 1

## your twilio account number
twilio_account: "XXXXXXXXXXX"

## your twilio account api key
twilio_key: "XXXXXXXXXXX"

## from is a string array of Twilio numbers to send from(you must own them)
## note: memeblast chooses which number to send from randomly
from:
  - "+1XXX8397357"
  - "+172XXX281921"
  - "+1720XXX1393"
  - "+172072XXX24"
  - "+17209XXX013"
  - "+172XXX91354"
  - "+1XXX7261600"
  - "+1720XXX5826"
  - "+17205XXX182"
  - "+17207XXX986"

## targets is a string array of numbers to send to. memeblast will send the same number of texts to each listed target.
targets:
  - "+130387XXX6"

## messages is a map of the text body to send : meme(img) url
messages:
  when the cats are away the mice will play: "http://1.bp.blogspot.com/_PG9h1CS1dfo/TAzK0FRyGpI/AAAAAAAAMYQ/yYo6lwDJ1o4/s400/ugly_cats_21.jpg"
```

Text received:

![](http://1.bp.blogspot.com/_PG9h1CS1dfo/TAzK0FRyGpI/AAAAAAAAMYQ/yYo6lwDJ1o4/s400/ugly_cats_21.jpg)

when the cats are away the mice will play



## Usage `memeblast`

```text
using config file: ./memeblast.yaml

------------------------------------------------------------
 __  __  ____  __  __  ____  ____  __      __    ___  ____ 
(  \/  )( ___)(  \/  )( ___)(  _ \(  )    /__\  / __)(_  _)
 )    (  )__)  )    (  )__)  ) _ < )(__  /(__)\ \__ \  )(  
(_/\/\_)(____)(_/\/\_)(____)(____/(____)(__)(__)(___/ (__) 
------------------------------------------------------------ 
Send sms blasts of memes from a randomized set of phonenumbers

(FOR EXPERIMENTATION PURPOSES ONLY)
(NOT FOR USE BY HUMANS)
SMS Laws: https://www.bulksms.com/resources/regulations/requirements-for-marketing-using-sms-messaging-in-the-usa.htm
------------------------------------------------------------

Usage:
  memeblast [command]

Available Commands:
  blast       start an sms blast (default: cat memes)
  debug       debug current config file
  help        Help about any command

Flags:
  -h, --help   help for memeblast

Use "memeblast [command] --help" for more information about a command.


```

