# SMSDOS

Annoy your friends by blasting them with silly sms messages

**FOR EXPERIMENTATION PURPOSES ONLY**

**NOT FOR USE BY HUMANS**

## Requirements
- a twilio account with active numbers
- friends that have phone numbers
- a config file named smsdos.yaml in $PWD 

## Example Config

```text
## number of randomized texts to send to each number(to)
send: 1

## from is a string array of Twilio numbers to send from(you must own them) 
## note: smsdos chooses which number to send from randomly
from:
  - "+17205835735"

## phone numbers is a string array of numbers to send to
to:
  - "+17703577402"
```


## Usage `smsdos`

```text
using config file: ./smsdos.yaml

 ______   _________   ______   _____    ______   ______ 
/ |      | | | | | \ / |      | | \ \  / |  | \ / |     
'------. | | | | | | '------. | |  | | | |  | | '------.
 ____|_/ |_| |_| |_|  ____|_/ |_|_/_/  \_|__|_/  ____|_/
                                                        
Annoy your friends by blasting them with silly sms messages

(FOR EXPERIMENTATION PURPOSES ONLY)
(NOT FOR USE BY HUMANS)
SMS Laws: https://www.bulksms.com/resources/regulations/requirements-for-marketing-using-sms-messaging-in-the-usa.htm

Usage:
  smsdos [command]

Available Commands:
  cats        start an cats sms blast
  debug       debug current config file
  help        Help about any command

Flags:
  -h, --help   help for smsdos

Use "smsdos [command] --help" for more information about a command.
```

