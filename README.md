# generate-fake-data

 🧬 CLI to generate fake data for testing

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/generate-fake-data)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/generate-fake-data/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/generate-fake-data.svg)](https://github.com/moul/generate-fake-data/releases)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/generate-fake-data.svg)](https://microbadger.com/images/moul/generate-fake-data)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/generate-fake-data/workflows/Go/badge.svg)](https://github.com/moul/generate-fake-data/actions?query=workflow%3AGo)
[![Release](https://github.com/moul/generate-fake-data/workflows/Release/badge.svg)](https://github.com/moul/generate-fake-data/actions?query=workflow%3ARelease)
[![PR](https://github.com/moul/generate-fake-data/workflows/PR/badge.svg)](https://github.com/moul/generate-fake-data/actions?query=workflow%3APR)
[![GolangCI](https://golangci.com/badges/github.com/moul/generate-fake-data.svg)](https://golangci.com/r/github.com/moul/generate-fake-data)
[![codecov](https://codecov.io/gh/moul/generate-fake-data/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/generate-fake-data)
[![Go Report Card](https://goreportcard.com/badge/moul.io/generate-fake-data)](https://goreportcard.com/report/moul.io/generate-fake-data)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/generate-fake-data/badge)](https://www.codefactor.io/repository/github/moul/generate-fake-data)


Fake data is based on https://github.com/brianvoe/gofakeit.

## Usage

`$> generate-fake-data -h`

[embedmd]:# (.tmp/usage.txt)
```txt
USAGE
  generate-fake-data

FLAGS
  -dict random      available: address, apache-log, beer, hacker, hipster, lorem-ipsum, phrase, question, quote, random, uuid
  -lines 42         amount of lines to generate
  -no-stderr false  disable writing to stderr (incompatible with no-stdout)
  -no-stdout false  disable writing to stdout (incompatible with no-stderr)
  -seed 0           seed random
  -sleep-max 300ms  maximum sleep duration between two lines
  -sleep-min 0s     minimum sleep duration between two lines
```

`$> generate-fake-data --dict=random`

[embedmd]:# (.tmp/random.txt)
```txt
635 Cove berg, Abbottfurt, Kansas 23215, Libyan Arab Jamahiriya
fefdbe0c-1028-47e1-80da-cd2d885f692c
"Flexitarian intelligentsia mumblecore lumbersexual occupy roof." - Taya Feil
80.219.187.139 - - [1989-10-30 16:54:12.195812429 +0000 UTC] 94282935 "PATCH et/blanditiis HTTP/1.1" 301 17 "-" ""Mozilla/5.0 (iPad; CPU OS 9_0_2 like Mac OS X; en-US) AppleWebKit/531.7.2 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.7.2""
Samuel Smith’s Oatmeal Stout
such is life
Samuel Smith’s Imperial IPA
396 Brook fort, Lake McLaughlin, North Carolina 17310, Armenia
882 North Loop view, Port McKenzie, Maryland 87318, Canada
54348 Inlet berg, South Nitzsche, Wyoming 39559, Hungary
```

`$> generate-fake-data --dict=address`

[embedmd]:# (.tmp/address.txt)
```txt
803 Harbor burgh, New Kreiger, New Jersey 58232, Central African Republic
45412 West Ridges ville, Anabelleland, Hawaii 84727, Uruguay
242 North Groves furt, Koeppstad, Minnesota 81674, Macao
6865 Circles berg, Gretafurt, California 43693, Nicaragua
638 Port Lights fort, Valerieshire, South Dakota 62848, Oman
878 Port Mountain berg, Schroedershire, Idaho 32083, Turks and Caicos Islands
8941 Villages chester, Ullrichmouth, South Carolina 59318, Cape Verde
6371 Valley chester, Daxhaven, Florida 20163, South Georgia and the South Sandwich Islands
39709 South Vista chester, Phoebeland, Massachusetts 69360, Gibraltar
7516 North Springs mouth, Hodkiewiczfurt, Virginia 29073, San Marino
```

`$> generate-fake-data --dict=address`

[embedmd]:# (.tmp/address.txt)
```txt
803 Harbor burgh, New Kreiger, New Jersey 58232, Central African Republic
45412 West Ridges ville, Anabelleland, Hawaii 84727, Uruguay
242 North Groves furt, Koeppstad, Minnesota 81674, Macao
6865 Circles berg, Gretafurt, California 43693, Nicaragua
638 Port Lights fort, Valerieshire, South Dakota 62848, Oman
878 Port Mountain berg, Schroedershire, Idaho 32083, Turks and Caicos Islands
8941 Villages chester, Ullrichmouth, South Carolina 59318, Cape Verde
6371 Valley chester, Daxhaven, Florida 20163, South Georgia and the South Sandwich Islands
39709 South Vista chester, Phoebeland, Massachusetts 69360, Gibraltar
7516 North Springs mouth, Hodkiewiczfurt, Virginia 29073, San Marino
```

`$> generate-fake-data --dict=apache-log`

[embedmd]:# (.tmp/apache-log.txt)
```txt
249.195.38.38 - - [1922-06-09 08:08:43.579751929 +0000 UTC] 68025971 "PUT facere/deleniti HTTP/1.1" 301 17 "-" ""Mozilla/5.0 (Windows; U; Windows NT 5.1) AppleWebKit/536.45.1 (KHTML, like Gecko) Version/5.0 Safari/536.45.1""
131.32.204.175 - - [1973-12-25 19:44:20.512849536 +0000 UTC] 28037524 "HEAD molestias/iste HTTP/1.1" 100 17 "-" ""Opera/9.88 (X11; Linux i686; en-US) Presto/2.11.218 Version/13.00""
149.192.105.11 - - [1922-03-20 07:46:47.97434687 +0000 UTC] 11972106 "PATCH facilis/voluptatem HTTP/1.1" 403 17 "-" ""Opera/10.98 (Windows NT 4.0; en-US) Presto/2.12.291 Version/12.00""
168.150.13.171 - - [1940-08-14 07:56:42.029503758 +0000 UTC] 37983735 "DELETE at/aut HTTP/1.1" 400 17 "-" ""Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_1 rv:7.0; en-US) AppleWebKit/531.26.1 (KHTML, like Gecko) Version/5.0 Safari/531.26.1""
253.193.178.45 - - [1988-06-15 14:41:34.824927687 +0000 UTC] 13890904 "HEAD quam/in HTTP/1.1" 204 17 "-" ""Mozilla/5.0 (Macintosh; PPC Mac OS X 10_7_8) AppleWebKit/5322 (KHTML, like Gecko) Chrome/37.0.878.0 Mobile Safari/5322""
127.167.179.32 - - [1902-04-09 17:48:31.904045112 +0000 UTC] 94859867 "PUT iusto/ratione HTTP/1.1" 200 17 "-" ""Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5342 (KHTML, like Gecko) Chrome/39.0.813.0 Mobile Safari/5342""
117.203.21.38 - - [1925-10-02 21:33:26.12022322 +0000 UTC] 92754440 "PATCH fugit/aut HTTP/1.1" 302 17 "-" ""Mozilla/5.0 (iPad; CPU OS 8_1_1 like Mac OS X; en-US) AppleWebKit/534.47.1 (KHTML, like Gecko) Version/4.0.5 Mobile/8B114 Safari/6534.47.1""
15.225.129.46 - - [1961-12-05 07:04:42.642391745 +0000 UTC] 87781661 "PATCH velit/eum HTTP/1.1" 200 17 "-" ""Opera/9.77 (X11; Linux i686; en-US) Presto/2.11.316 Version/10.00""
12.171.201.129 - - [1979-08-15 04:51:32.664630958 +0000 UTC] 82133365 "DELETE error/quas HTTP/1.1" 100 17 "-" ""Mozilla/5.0 (Windows 95) AppleWebKit/5350 (KHTML, like Gecko) Chrome/36.0.882.0 Mobile Safari/5350""
178.105.136.112 - - [1980-04-04 00:23:11.295759587 +0000 UTC] 2014487 "GET iste/optio HTTP/1.1" 405 17 "-" ""Mozilla/5.0 (Windows; U; Windows NT 5.1) AppleWebKit/532.43.5 (KHTML, like Gecko) Version/4.2 Safari/532.43.5""
```

`$> generate-fake-data --dict=beer`

[embedmd]:# (.tmp/beer.txt)
```txt
St. Bernardus Abt 12
Samuel Smith’s Oatmeal Stout
Schneider Aventinus
Pliny The Elder
Chimay Grande Réserve
Racer 5 India Pale Ale, Bear Republic Bre
Weihenstephaner Hefeweissbier
Edmund Fitzgerald Porter
Oak Aged Yeti Imperial Stout
Sierra Nevada Celebration Ale
```

`$> generate-fake-data --dict=hacker`

[embedmd]:# (.tmp/hacker.txt)
```txt
We need to parse the neural AGP card!
We need to program the online JBOD transmitter!
We need to calculate the digital SMTP bus!
The HTTP monitor is down, parse the back-end system so we can navigate the SMS panel!
If we parse the bandwidth, we can get to the SAS transmitter through the bluetooth SSL pixel!
You can't back up the program without quantifying the auxiliary PCI firewall!
You can't copy the transmitter without backing up the online AI feed!
Try to back up the PCI alarm, maybe it will reboot the open-source program!
I'Ll connect the auxiliary SMS interface, that should synthesize the PNG feed!
Try to override the SQL alarm, maybe it will bypass the neural array!
```

`$> generate-fake-data --dict=hipster`

[embedmd]:# (.tmp/hipster.txt)
```txt
Gastropub five dollar toast mumblecore kinfolk loko bushwick semiotics pickled.
Tote bag neutra loko tote bag +1 salvia organic umami.
Polaroid synth wolf roof PBR&B asymmetrical craft beer taxidermy.
Everyday venmo shabby chic art party marfa flexitarian intelligentsia mumblecore.
Lumbersexual occupy roof you probably haven't heard of them kitsch wayfarers hashtag pitchfork.
Shabby chic echo cronut gentrify wolf jean shorts pabst normcore.
XOXO DIY 3 wolf moon plaid roof fanny pack Thundercats typewriter.
Deep v literally lo-fi cardigan hammock umami keytar polaroid.
Schlitz yr YOLO sriracha small batch neutra truffaut sustainable.
Slow-carb fanny pack neutra wayfarers helvetica polaroid vinegar biodiesel.
```

`$> generate-fake-data --dict=lorem-ipsum`

[embedmd]:# (.tmp/lorem-ipsum.txt)
```txt
At illum ut est sit soluta nulla numquam.
Nobis sunt quaerat ea dolores facere deleniti culpa.
Numquam ut distinctio maxime consequatur est qui corporis.
Sunt officia odit et quia odit molestias voluptas.
Porro repellendus magnam ipsa corporis eos rem non.
Hic esse optio quisquam hic natus earum molestias.
Iste architecto porro et blanditiis iste eum repellendus.
Nostrum qui eius suscipit fugit quia quo et.
Nesciunt quod fuga ut vel pariatur libero sequi.
Rerum omnis soluta facilis voluptatem possimus et voluptas.
```

`$> generate-fake-data --dict=phrase`

[embedmd]:# (.tmp/phrase.txt)
```txt
Lord knows
when is closing time
many thanks
could fit on the back of a postage stamp
size matters
could have, would have, should have
I'm tired
I'm worried
ultra vires
you win
```

`$> generate-fake-data --dict=question`

[embedmd]:# (.tmp/question.txt)
```txt
Five dollar toast mumblecore kinfolk loko?
Semiotics pickled tote bag neutra?
Tote bag +1 salvia organic?
Polaroid synth wolf roof PBR&B asymmetrical craft beer taxidermy?
Venmo shabby chic art party?
Flexitarian intelligentsia mumblecore lumbersexual occupy roof?
Kitsch wayfarers hashtag pitchfork shabby chic?
Cronut gentrify wolf jean shorts pabst normcore?
DIY 3 wolf moon plaid roof fanny pack Thundercats typewriter deep v?
Lo-fi cardigan hammock umami?
```

`$> generate-fake-data --dict=quote`

[embedmd]:# (.tmp/quote.txt)
```txt
"Five dollar toast mumblecore kinfolk loko." - Stuart Pacocha
"Tote bag neutra loko." - Jaylon Johns
"Organic umami polaroid synth wolf roof PBR&B." - Brendon Jast
"Everyday venmo shabby chic art party marfa flexitarian." - Ryann Johns
"Occupy roof you probably haven't heard of them kitsch wayfarers hashtag." - Leonie Goodwin
"Cronut gentrify wolf jean shorts pabst normcore." - Darion Lehner
"Plaid roof fanny pack Thundercats typewriter." - Hubert Dare
"Cardigan hammock umami keytar polaroid." - Uriah Erdman
"Sriracha small batch neutra." - June Terry
"Fanny pack neutra wayfarers helvetica polaroid vinegar biodiesel chambray street." - Caden Parker
```

`$> generate-fake-data --dict=uuid`

[embedmd]:# (.tmp/uuid.txt)
```txt
538c7f96-b164-4f1b-97bb-9f4bb472e89f
5b1484f2-5209-49d9-b43e-92ba09dd9d52
dfd79b4d-7642-4b61-ba0c-9f9f0d3ba55b
0cc0d614-4c88-4535-841a-cbe0709b0758
083f61d3-75bc-42b4-9df4-f91929e18fda
9e6f82e5-4e74-4e81-a79e-4bbd6fe34cdc
ba843ee8-d63e-4c4f-be1c-ebea546d8fac
13dd1aac-04ce-4ea2-877c-5579cfa2c78e
1b0bafae-881b-42a7-9110-8a42ed3c903c
aa43465a-7862-4616-978a-ed0ce3c6c4f3
```

## Install

### Using go

```console
$ go get -u moul.io/generate-fake-data
```

### Releases

See https://github.com/moul/generate-fake-data/releases

## Contribute

![Contribute <3](https://raw.githubusercontent.com/moul/moul/master/contribute.gif)

I really welcome contributions. Your input is the most precious material. I'm well aware of that and I thank you in advance. Everyone is encouraged to look at what they can do on their own scale; no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./CONTRIBUTING.md)

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/generate-fake-data/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/generate-fake-data/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/generate-fake-data/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/generate-fake-data.svg)](https://starchart.cc/moul/generate-fake-data)

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
