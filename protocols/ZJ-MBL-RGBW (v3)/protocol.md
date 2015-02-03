# The droids we are looking for

<img src="sample.png" alt="Subject appearance" width="100%">

This document describes GATT characteristics used to control Zengge smart RGBW BLE bulbs (like on the picture above). All the text below is a result of my own experience of reverse-engineering its protocol. Use it at your own risk, **I am not responsible for your hardware.**

Looking through the reverse engineered code I figured out that there are potentially two slightly different kinds of bulbs. They differ by names and versions. Bulbs whose names start with `LEDBlue` or `LEDBLE` belong to the first group, not subject of this doc. Second group of devices is described here. Also, some commands differ depending on the device version (can be known [by querying it's status](#device_version)). The commands below are tested on `LEDnet-...`-named bulb of third version.

# Communication principles

There are three main ways to interact with a bulb:

1. **Fire and forget.** Only writes are supported. No response or acknowledgement.
1. **Write and listen** for notifications. Write a characteristic while subscribed for notifications. As a result of value write, another characteristics may fire a notification.
1. **Direct read of a characteristic**. Read-only access to some parameters.

Most of the communication with a bulb is done via two characteristcs — `FFE9` and `FFE4` — in a write-and-listen manner: you write a data into `FFE9` and listen for `FFE4`'s notifications. 

# The protocol

## Status

Status request is used to query some generic bulb parameters, like power status, device version, color or predefined mode id. Status query is done in "write and listen" manner, so it will not work in non-interactive `gatttool` mode (you won't be able to see the result).

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write and listen</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFE9</code></td>
		</tr>
		<tr>
			<td>Notification from</td>
			<td><code>FFE4</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td>
				Constant, <code>[0xEF, 0x01, 0x77]</code>
			</td>
		</tr>
		<tr>
			<td>Notification</td>
			<td>See the description below</td>
		</tr>
	</tbody>
</table>

#### Notification description

Resulting notification _must_ be 12 bytes long.

1. `result[0]` _must_ be equal to magic constant `0x66`
1. `result[1]`: ???
1. `result[2]`: power status:
	+ `0x23` is for "ON".
	+ `0x24` is for "OFF".
1. `result[3]`: mode:
	+ `0x25`-`0x38`: [build-in mode](#built-in-modes)
	+ `0x41`: static color mode
1. `result[4]`: ???
1. `result[5]`: [speed](#speed)
1. `result[6]`: red color component
1. `result[7]`: green color component
1. `result[8]`: blue color component
1. `result[9]`: ???
1. `result[10]`: <a name="device_version"></a>device version
1. `result[11]` _must_ be equal to magic constant `0x99`

#### Examples

<details>

<table>
	<thead>
		<tr>
			<th>Magic</th>
			<th>???</th>
			<th>Power</th>
			<th>Mode</th>
			<th>???</th>
			<th>Speed</th>
			<th>R</th>
			<th>G</th>
			<th>B</th>
			<th>???</th>
			<th>Version</th>
			<th>Magic</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x14</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Static color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x21</code></td>
			<td title="Speed (does not matter in static color mode)"><code>0x1F</code></td>
			<td title="Red color component"><code>0xFF</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0x00</code></td>
			<td title="Unknown"><code>0x00</code></td>
			<td title="Device version"><code>0x03</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Static&nbsp;red&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x14</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Static color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x21</code></td>
			<td title="Speed (does not matter in static color mode)"><code>0x1F</code></td>
			<td title="Red color component"><code>0x00</code></td>
			<td title="Green color component"><code>0xFF</code></td>
			<td title="Blue color component"><code>0x00</code></td>
			<td title="Unknown"><code>0x00</code></td>
			<td title="Device version"><code>0x03</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Static&nbsp;green&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x14</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Static color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x21</code></td>
			<td title="Speed (does not matter in static color mode)"><code>0x1F</code></td>
			<td title="Red color component"><code>0x00</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0xFF</code></td>
			<td title="Unknown"><code>0x00</code></td>
			<td title="Device version"><code>0x03</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Static&nbsp;blue&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x14</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Static color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x21</code></td>
			<td title="Speed (does not matter in static color mode)"><code>0x1F</code></td>
			<td title="Red color component"><code>0x5A</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0x9D</code></td>
			<td title="Unknown"><code>0x00</code></td>
			<td title="Device version"><code>0x03</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Static&nbsp;violet&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x14</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Built-in mode"><code>0x27</code></td>
			<td title="Unknown"><code>0x21</code></td>
			<td title="The slowest possible"><code>0x1F</code></td>
			<td title="Red color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Green color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Blue color component (does not matter in built-in mode)"><code>0xFF</code></td>
			<td title="Unknown"><code>0x00</code></td>
			<td title="Device version"><code>0x03</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Built&#8209;in&nbsp;mode&nbsp;<code>0x27</code>&nbsp;at&nbsp;speed&nbsp;<code>0x1F</code>&nbsp;(the&nbsp;slowest&nbsp;possible)</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x14</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Built-in mode"><code>0x34</code></td>
			<td title="Unknown"><code>0x21</code></td>
			<td title="Fast, but not the fastest"><code>0x10</code></td>
			<td title="Red color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Green color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Blue color component (does not matter in built-in mode)"><code>0xFF</code></td>
			<td title="Unknown"><code>0x00</code></td>
			<td title="Device version"><code>0x03</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Built&#8209;in&nbsp;mode&nbsp;<code>0x34</code>&nbsp;at&nbsp;speed&nbsp;<code>0x10</code>&nbsp;(fast)</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x14</code></td>
			<td title="LEDs are OFF"><code>0x24</code></td>
			<td title="Built-in mode (does not matter when OFF)"><code>0x34</code></td>
			<td title="Unknown"><code>0x21</code></td>
			<td title="Speed (does not matter when OFF)"><code>0x10</code></td>
			<td title="Red color component (does not matter when OFF)"><code>0x00</code></td>
			<td title="Green color component (does not matter when OFF)"><code>0x00</code></td>
			<td title="Blue color component (does not matter when OFF)"><code>0xFE</code></td>
			<td title="Unknown"><code>0x00</code></td>
			<td title="Device version"><code>0x03</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Turned&nbsp;off&nbsp;built&#8209;in&nbsp;mode&nbsp;<code>0x34</code>&nbsp;at&nbsp;speed&nbsp;<code>0x10</code>&nbsp;(fast)</td>
		</tr>
	</tbody>
</table>

</details>

## Power

### Query for current status

Power status can be known from `FFF3` characteristic.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Direct read</td>
		</tr>
		<tr>
			<td>Read from</td>
			<td><code>FFF3</code></td>
		</tr>
		<tr>
			<td>Read result</td>
			<td>
				1. <code>0x3B</code> when the bulb is off<br/>
				2. <code>0xFF</code> when the bulb is on
			</td>
		</tr>
	</tbody>
</table>

### Set current status

Power status can be set via `FFF1` and `FFF2` characteristics. It's a two step process: first you write `0x04` in `FFF1` and then you set power status in `FFF2`. The requests are done with 200ms interval.

#### Requests

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFF1</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td><code>0x04</code></td>
		</tr>
	</tbody>
</table>

…wait 200ms…

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFF2</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td>
				1. <code>0x00</code> turn off<br/>
				2. <code>0x3F</code> turn on
			</td>
		</tr>
	</tbody>
</table>

#### Example

<details>

<table>
	<thead>
		<tr>
			<th>Request</th>
			<th>Action</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>
				<pre>gatttool -b B4:99:4C:2A:0E:4A --char-write-req -a 0x0017 -n 04 &&
sleep 0.2s &&
gatttool -b B4:99:4C:2A:0E:4A --char-write-req -a 0x001a -n 00
				</pre>
			</td>
			<td>Turn&nbsp;the&nbsp;bulb&nbsp;off</td>
		</tr>
		<tr>
			<td>
				<pre>gatttool -b B4:99:4C:2A:0E:4A --char-write-req -a 0x0017 -n 04 &&
sleep 0.2s &&
gatttool -b B4:99:4C:2A:0E:4A --char-write-req -a 0x001a -n FF
				</pre>
			</td>
			<td>Turn&nbsp;the&nbsp;bulb&nbsp;on</td>
		</tr>
	</tbody>
</table>

</details>

## Static color mode

Static color mode is set via write request to `FFE9` characteristic.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFE9</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td>See below</td>
		</tr>
	</tbody>
</table>

#### Payload description

Payload _must_ be 7 bytes long.

1. `payload[0]` _must_ be equal to magic constant `0x56`
1. `payload[1]`: red color component
1. `payload[2]`: green color component
1. `payload[3]`: blue color component
1. `payload[4]` _must_ be equal to magic constant `0x00`
1. `payload[5]` _must_ be equal to magic constant `0xF0`
1. `payload[6]` _must_ be equal to magic constant `0xAA`

#### Examples

<details>

<table>
	<thead>
		<tr>
			<th>Magic</th>
			<th>R</th>
			<th>G</th>
			<th>B</th>
			<th>Magic</th>
			<th>Magic</th>
			<th>Magic</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Constant value 0x56"><code>0x56</code></td>
			<td title="Red color component"><code>0xFF</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0x00</code></td>
			<td title="Constant value 0x00"><code>0x00</code></td>
			<td title="Constant value 0xF0"><code>0xF0</code></td>
			<td title="Constant value 0xAA"><code>0xAA</code></td>
			<td>Static&nbsp;red&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x56"><code>0x56</code></td>
			<td title="Red color component"><code>0x00</code></td>
			<td title="Green color component"><code>0xFF</code></td>
			<td title="Blue color component"><code>0x00</code></td>
			<td title="Constant value 0x00"><code>0x00</code></td>
			<td title="Constant value 0xF0"><code>0xF0</code></td>
			<td title="Constant value 0xAA"><code>0xAA</code></td>
			<td>Static&nbsp;green&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x56"><code>0x56</code></td>
			<td title="Red color component"><code>0x00</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0xFF</code></td>
			<td title="Constant value 0x00"><code>0x00</code></td>
			<td title="Constant value 0xF0"><code>0xF0</code></td>
			<td title="Constant value 0xAA"><code>0xAA</code></td>
			<td>Static&nbsp;blue&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x56"><code>0x56</code></td>
			<td title="Red color component"><code>0x5A</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0x9D</code></td>
			<td title="Constant value 0x00"><code>0x00</code></td>
			<td title="Constant value 0xF0"><code>0xF0</code></td>
			<td title="Constant value 0xAA"><code>0xAA</code></td>
			<td>Static&nbsp;violet&nbsp;color</td>
		</tr>
	</tbody>
</table>

</details>

## Built-in mode

Built-in mode is set via write request to `FFE9` characteristic.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFE9</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td>See below</td>
		</tr>
	</tbody>
</table>

#### Payload description

Payload _must_ be 4 bytes long.

1. `payload[0]` _must_ be equal to magic constant `0xBB`
1. `payload[1]`: [build-in mode](#built-in-modes)
1. `payload[2]`: [speed](#speed)
1. `payload[3]` _must_ be equal to magic constant `0x44`

#### Examples

<details>

<table>
	<thead>
		<tr>
			<th>Magic</th>
			<th>Mode</th>
			<th>Speed</th>
			<th>Magic</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Constant value 0xBB"><code>0xBB</code></td>
			<td title="Green gradual change"><code>0x27</code></td>
			<td title="The slowest possible"><code>0x1F</code></td>
			<td title="Constant value 0x44"><code>0x44</code></td>
			<td>Built&#8209;in&nbsp;mode&nbsp;<code>0x27</code>&nbsp;at&nbsp;speed&nbsp;<code>0x1F</code>&nbsp;(the&nbsp;slowest&nbsp;possible)</td>
		</tr>
		<tr>
			<td title="Constant value 0xBB"><code>0xBB</code></td>
			<td title="Yellow strobe flash"><code>0x34</code></td>
			<td title="Fast"><code>0x10</code></td>
			<td title="Constant value 0x44"><code>0x44</code></td>
			<td>Built&#8209;in&nbsp;mode&nbsp;<code>0x34</code>&nbsp;at&nbsp;speed&nbsp;<code>0x10</code>&nbsp;(fast)</td>
		</tr>
	</tbody>
</table>

</details>

## Clock

### Query for current clock value

Current clock can be read from `FE01` characteristic. Epoch start for the bulb is `2000-01-01 00:00:00`; every time you turn the electricity off (on a hardware level, not like described above) the clock is reset to that value.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Direct read</td>
		</tr>
		<tr>
			<td>Read from</td>
			<td><code>FE01</code></td>
		</tr>
		<tr>
			<td>Read result</td>
			<td>See below</td>
		</tr>
	</tbody>
</table>

#### Response

The response is 7 bytes long.

1. `result[0]`: seconds
1. `result[1]`: minutes
1. `result[2]`: hours (24 hours format)
1. `result[3]`: day of month, starting from 1
1. `result[4]`: month (1 is Jan., 2 is Feb., etc)
1. `result[5]`: lower byte of the year
1. `result[6]`: upper byte of the year

#### Example

<details>

	gatttool -b B4:99:4C:2A:0E:4A --char-read -a 0x0086
	Characteristic value/descriptor: 08 36 01 01 01 df 07

<table>
	<thead>
		<tr>
			<th>Seconds</th>
			<th>Minutes</th>
			<th>Hours</th>
			<th>Date</th>
			<th>Month</th>
			<th>Lower year</th>
			<th>Upper year</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Seconds (08)"><code>0x08</code></td>
			<td title="Minutes (54)"><code>0x36</code></td>
			<td title="Hours (01)"><code>0x01</code></td>
			<td title="Date (01)"><code>0x01</code></td>
			<td title="Month (January)"><code>0x01</code></td>
			<td title="Lower byte of year (2015 & 0xFF == 223)"><code>0xDF</code></td>
			<td title="Upper byte of year (2015 >> 8 == 7)"><code>0x07</code></td>
			<td>01/01/2015&nbsp;01:54:08</td>
		</tr>
	</tbody>
</table>

</details>

### Set clock value

Current clock can be set by writing `FE01` characteristic (the same that is used to query clock).

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FE01</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td>See below</td>
		</tr>
	</tbody>
</table>

#### Payload description

Payload _must_ be 7 bytes long.

1. `result[0]`: seconds
1. `result[1]`: minutes
1. `result[2]`: hours (24 hours format)
1. `result[3]`: day of month, starting from 1
1. `result[4]`: month (1 is Jan., 2 is Feb., etc)
1. `result[5]`: lower byte of the year
1. `result[6]`: upper byte of the year

#### Example

<details>

	gatttool -b B4:99:4C:2A:0E:4A --char-write -a 0x0086 -n 3b1f011e01df07

<table>
	<thead>
		<tr>
			<th>Seconds</th>
			<th>Minutes</th>
			<th>Hours</th>
			<th>Date</th>
			<th>Month</th>
			<th>Lower year</th>
			<th>Upper year</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Seconds (59)"><code>0x3b</code></td>
			<td title="Minutes (31)"><code>0x1F</code></td>
			<td title="Hours (01)"><code>0x01</code></td>
			<td title="Date (30)"><code>0x1E</code></td>
			<td title="Month (January)"><code>0x01</code></td>
			<td title="Lower byte of year (2015 & 0xFF == 223)"><code>0xDF</code></td>
			<td title="Upper byte of year (2015 >> 8 == 7)"><code>0x07</code></td>
			<td>01/30/2015&nbsp;01:31:59</td>
		</tr>
	</tbody>
</table>

</details>

# Magic constants

## Built-in modes

1. `0x25`: Seven color cross fade
1. `0x26`: Red gradual change
1. `0x27`: Green gradual change
1. `0x28`: Blue gradual change
1. `0x29`: Yellow gradual change
1. `0x2a`: Cyan gradual change
1. `0x2b`: Purple gradual change
1. `0x2c`: White gradual change
1. `0x2d`: Red, Green cross fade
1. `0x2e`: Red blue cross fade
1. `0x2f`: Green blue cross fade
1. `0x30`: Seven color stobe flash
1. `0x31`: Red strobe flash
1. `0x32`: Green strobe flash
1. `0x33`: Blue strobe flash
1. `0x34`: Yellow strobe flash
1. `0x35`: Cyan strobe flash
1. `0x36`: Purple strobe flash
1. `0x37`: White strobe flash
1. `0x38`: Seven color jumping change


## Speed

Some operational modes take a speed parameter that controls how fast the colors are changed. `0x01` is the fastest, `0x1F` is the slowest.
