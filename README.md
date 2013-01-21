fan3xxnsa
=========

Fan control daemon for NSA 310 NAS server.

This is very simple and lightweight daemon that reads case temperature of Zyxel
NSA-310 and sets fan speed according to the temperature.

How does it work
----------------

The NSA 310 provides three temperatures. The first two are most likely case
temperatures, because according to all my measurements, they were always the
same (+- one degree) as the disc temperature reported by SMART. The third
temperature is from the CPU.

Please note the daemon only reads the case temperature (temp1), it ignores CPU
or HDD temperatures. Because it does not make any sense to measure it (and also
SMART spins up disc). Fan speed depends only on temp1, nothing else.

I use WD GreenPower in my NAS and the disc is 42 degrees Celsius when idle or
light work, so the threasold is 43 degrees. Also I have noticed this cheap and
small fan creates noise when rotating bellow 140 value, so the lowest is 160.
Make sure screws in the back side are loose, do not tighen them much as it can
put a pressure on the fan and create even more noise.

This is how it sets pwm according to temperature:

	Temp: 41 Pwm: 0
	Temp: 42 Pwm: 0
	Temp: 43 Pwm: 160
	Temp: 44 Pwm: 165
	Temp: 45 Pwm: 170
	Temp: 46 Pwm: 175
	Temp: 47 Pwm: 180
	Temp: 48 Pwm: 185
	Temp: 49 Pwm: 190
	Temp: 50 Pwm: 195
	Temp: 51 Pwm: 200
	Temp: 52 Pwm: 205
	Temp: 53 Pwm: 210
	Temp: 54 Pwm: 215
	Temp: 55 Pwm: 220
	Temp: 56 Pwm: 225
	Temp: 57 Pwm: 230
	Temp: 58 Pwm: 235
	Temp: 59 Pwm: 240
	Temp: 60 Pwm: 245
	Temp: 61 Pwm: 250
	Temp: 62 Pwm: 255
	Temp: 63 Pwm: 255

How to install
--------------

Very simple. First, you need Go language (golang). Please note Go 1.0 WILL NOT
compile on ARM5, you either need some development snapshot or you can checkout
master (called "default") branch. Then install and set GOROOT and PATH
variables.

To install fan3xxnsa, just do:

	git clone
	go build

And copy the binary file to a destination directory. If you are fine with ~/bin
there is a script install.sh that does this for you (and also strips debug
information from the binary).

Before you start the daemon, you need to enable fan control with:

	echo 1 > /sys/class/i2c-dev/i2c-0/device/0-002e/pwm1_enable

Then you can start it on background using nohup for example:

	nohup /root/bin/fan3xxnsa 1>/dev/null 2>&1 &

Must be executed as root, or you can change permissions on the pwm1 file.

Hacking
-------

The code is very easy to understand, feel free to send patches. If you want to
play with fans and temperatures, go to /sys/class/i2c-dev/i2c-0/device/0-002e
directory (it can be different according to your NSA model or kernel version -
try to find "pwm1_enable" file in the /sys/class directory).

Patches for other NSA models or kernel versions appreciated.
