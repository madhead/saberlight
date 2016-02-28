# The droids we are looking for

<img src="../ZJ-MBL-RGBW (v3)/sample.png" alt="Subject appearance" width="100%">

This document describes GATT characteristics used to control Triones smart RGBW BLE bulbs (like on the picture above). Yes, they look the same as [Zengge bulbs](../ZJ-MBL-RGBW (v3)/protocol.md). My guess is that all cheap Chinese bulbs looks identically to each other, and only differ in GATT characteristics UUIDs.

All the text below is a result of my own experience of reverse-engineering its protocol. Use it at your own risk, **I am not responsible for your hardware.**

# Communication principles

TODO

# The protocol

## Status

Status request is used to query some generic bulb parameters, like power status, current color or built-in mode. Status query is done in "write and listen" manner, so it will not work in non-interactive `gatttool` mode (you won't be able to see the result). Check [the code](../../app/commands/status.go) for more details.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write and listen</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFD9</code></td>
		</tr>
		<tr>
			<td>Notification from</td>
			<td><code>FFD4</code></td>
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
1. `result[9]`: white color intensity (when the bulb is in white color mode)
1. `result[10]`: ???
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
			<th>W</th>
			<th>???</th>
			<th>Magic</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x15</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Static color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x20</code></td>
			<td title="Speed (does not matter in static color mode)"><code>0x00</code></td>
			<td title="Red color component"><code>0xFF</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0x00</code></td>
			<td title="White color intensity (does not matter in static color mode)"><code>0x00</code></td>
			<td title="Unknown"><code>0x06</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Static&nbsp;red&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x15</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Static color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x20</code></td>
			<td title="Speed (does not matter in static color mode)"><code>0x00</code></td>
			<td title="Red color component"><code>0x00</code></td>
			<td title="Green color component"><code>0xFF</code></td>
			<td title="Blue color component"><code>0x00</code></td>
			<td title="White color intensity (does not matter in static color mode)"><code>0x00</code></td>
			<td title="Unknown"><code>0x06</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Static&nbsp;green&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x15</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Static color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x20</code></td>
			<td title="Speed (does not matter in static color mode)"><code>0x00</code></td>
			<td title="Red color component"><code>0x00</code></td>
			<td title="Green color component"><code>0x00</code></td>
			<td title="Blue color component"><code>0xFF</code></td>
			<td title="White color intensity (does not matter in static color mode)"><code>0x00</code></td>
			<td title="Unknown"><code>0x06</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Static&nbsp;blue&nbsp;color</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x15</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="White color mode"><code>0x41</code></td>
			<td title="Unknown"><code>0x20</code></td>
			<td title="Speed (does not matter in white color mode)"><code>0x00</code></td>
			<td title="Red color component (does not matter in white color mode)"><code>0x00</code></td>
			<td title="Green color component (does not matter in white color mode)"><code>0x00</code></td>
			<td title="Blue color component (does not matter in white color mode)"><code>0x00</code></td>
			<td title="White color intensity"><code>0x30</code></td>
			<td title="Unknown"><code>0x06</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>White&nbsp;color&nbsp;with&nbsp;low&nbsp;intensity</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x15</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Built-in mode"><code>0x27</code></td>
			<td title="Unknown"><code>0x20</code></td>
			<td title="Slowest possible speed"><code>0x1F</code></td>
			<td title="Red color component (does not matter in built-in mode)"><code>0xFF</code></td>
			<td title="Green color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Blue color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="White color intensity (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Unknown"><code>0x06</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Built&#8209;in&nbsp;mode&nbsp;<code>0x27</code>&nbsp;at&nbsp;speed&nbsp;<code>0x1F</code>&nbsp;(the&nbsp;slowest&nbsp;possible)</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x15</code></td>
			<td title="LEDs are ON"><code>0x23</code></td>
			<td title="Built-in mode"><code>0x34</code></td>
			<td title="Unknown"><code>0x20</code></td>
			<td title="Pretty fast speed"><code>0x10</code></td>
			<td title="Red color component (does not matter in built-in mode)"><code>0xFF</code></td>
			<td title="Green color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Blue color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="White color intensity (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Unknown"><code>0x06</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Built&#8209;in&nbsp;mode&nbsp;<code>0x34</code>&nbsp;at&nbsp;speed&nbsp;<code>0x10</code>&nbsp;(fast)</td>
		</tr>
		<tr>
			<td title="Constant value 0x66"><code>0x66</code></td>
			<td title="Unknown"><code>0x15</code></td>
			<td title="LEDs are OFF"><code>0x24</code></td>
			<td title="Built-in mode"><code>0x34</code></td>
			<td title="Unknown"><code>0x20</code></td>
			<td title="Pretty fast speed"><code>0x10</code></td>
			<td title="Red color component (does not matter in built-in mode)"><code>0xFF</code></td>
			<td title="Green color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Blue color component (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="White color intensity (does not matter in built-in mode)"><code>0x00</code></td>
			<td title="Unknown"><code>0x06</code></td>
			<td title="Constant value 0x99"><code>0x99</code></td>
			<td>Bulb&nbsp;is&nbsp;turned&nbsp;off&nbsp;with&nbsp;built&#8209;in&nbsp;mode&nbsp;<code>0x34</code>&nbsp;at&nbsp;speed&nbsp;<code>0x10</code>&nbsp;(fast)</td>
		</tr>
	</tbody>
</table>

</details>

## Power

### Set current power status

Power is turned on and off via write request to `FFD9` characteristic under `FFD5` servce. Check [the code](../../app/commands/power.go) for more details.

#### Requests

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFD9</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td>See below</td>
		</tr>
	</tbody>
</table>

#### Payload description

Payload _must_ be 3 bytes long.

1. `payload[0]` _must_ be equal to magic constant `0xCC`
1. `payload[3]`: `0x23` for "ON" and `0x24` for "OFF"
1. `payload[6]` _must_ be equal to magic constant `0x33`

#### Example

<details>
<table>
	<thead>
		<tr>
			<th>Magic</th>
			<th>Power status</th>
			<th>Magic</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Constant value 0xCC"><code>0xCC</code></td>
			<td title="ON"><code>0x23</code></td>
			<td title="Constant value 0x33"><code>0x33</code></td>
			<td>Turn power on</td>
		</tr>
		<tr>
			<td title="Constant value 0xCC"><code>0xCC</code></td>
			<td title="OFF"><code>0x24</code></td>
			<td title="Constant value 0x33"><code>0x33</code></td>
			<td>Turn power off</td>
		</tr>
	</tbody>
</table>
</details>

## Static color mode

Static color mode is set via write request to `FFD9` characteristic under `FFD5` servce. Check [the code](../../app/commands/color.go) for more details.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFD9</code></td>
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

## White color

White color is set via write request to `FFD9` characteristic under `FFD5` servce. Check [the code](../../app/commands/color.go) for more details.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFD9</code></td>
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
1. `payload[1]`: not used
1. `payload[2]`: not used
1. `payload[3]`: not used
1. `payload[4]`: intensity
1. `payload[5]` _must_ be equal to magic constant `0x0F`
1. `payload[6]` _must_ be equal to magic constant `0xAA`

#### Examples

<details>
<table>
	<thead>
		<tr>
			<th>Magic</th>
			<th>N/A</th>
			<th>N/A</th>
			<th>N/A</th>
			<th>Intensity</th>
			<th>Magic</th>
			<th>Magic</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Constant value 0x56"><code>0x56</code></td>
			<td title="N/A"><code>0xDE</code></td>
			<td title="N/A"><code>0xAD</code></td>
			<td title="N/A"><code>0xFF</code></td>
			<td title="Intensity"><code>0x01</code></td>
			<td title="Constant value 0x0F"><code>0x0F</code></td>
			<td title="Constant value 0xAA"><code>0xAA</code></td>
			<td>Lowest possible intensity</td>
		</tr>
		<tr>
			<td title="Constant value 0x56"><code>0x56</code></td>
			<td title="N/A"><code>0xCA</code></td>
			<td title="N/A"><code>0xFE</code></td>
			<td title="N/A"><code>0x00</code></td>
			<td title="Intensity"><code>0xFF</code></td>
			<td title="Constant value 0x0F"><code>0x0F</code></td>
			<td title="Constant value 0xAA"><code>0xAA</code></td>
			<td>Highest possible intensity</td>
		</tr>
	</tbody>
</table>
</details>

## Built-in mode

Built-in mode is set via write request to `FFD9` characteristic. Check [the code](../../app/commands/mode.go) for more details.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFD9</code></td>
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

Unfortunately, I do not know the way to query for current bulb's clock value.

### Set clock value

Current clock can be set by writing `FFD9` characteristic. Check [the code](../../app/commands/time.go) for more details.

#### Request

<table>
	<tbody>
		<tr>
			<td>Type</td>
			<td>Write</td>
		</tr>
		<tr>
			<td>Write to</td>
			<td><code>FFD9</code></td>
		</tr>
		<tr>
			<td>Payload</td>
			<td>See below</td>
		</tr>
	</tbody>
</table>

#### Payload description

Payload _must_ be 11 bytes long.

1. `payload[0]` _must_ be equal to magic constant `0x10`
1. `payload[1]` year divided by 100
1. `payload[2]` remainder of dividing year by 100
1. `payload[3]` month (1 is Jan., 2 is Feb., etc)
1. `payload[4]` day of month, starting from 1
1. `payload[5]` hours (24 hours format)
1. `payload[6]` minutes
1. `payload[7]` seconds
1. `payload[8]` day of week (SUN is 0)
1. `payload[9]` _must_ be equal to magic constant `0x00`
1. `payload[10]` _must_ be equal to magic constant `0x01`

#### Example

<details>
<table>
	<thead>
		<tr>
			<th>Magic</th>
			<th>Upper year</th>
			<th>Lower year</th>
			<th>Month</th>
			<th>Date</th>
			<th>Hours</th>
			<th>Minutes</th>
			<th>Seconds</th>
			<th>Day of week</th>
			<th>Magic</th>
			<th>Magic</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td title="Constant value 0x10"><code>0x10</code></td>
			<td title="Upper byte of year (2016 / 100 == 20)"><code>0x14</code></td>
			<td title="Lower byte of year (2016 % 100 == 16)"><code>0x10</code></td>
			<td title="Month (February)"><code>0x02</code></td>
			<td title="Date (28)"><code>0x1C</code></td>
			<td title="Hours (05)"><code>0x05</code></td>
			<td title="Minutes (07)"><code>0x07</code></td>
			<td title="Seconds (24)"><code>0x18</code></td>
			<td title="Day of week (Sunday, 7)"><code>0x07</code></td>
			<td title="Constant value 0x00"><code>0x00</code></td>
			<td title="Constant value 0x01"><code>0x01</code></td>
			<td>Sun Feb 28 05:07:24 2016</td>
		</tr>
	</tbody>
</table>
</details>

## Timings

TODO

# Magic constants

## Built-in modes

1. `0x25`: Seven color cross fade
1. `0x26`: Red gradual change
1. `0x27`: Green gradual change
1. `0x28`: Blue gradual change
1. `0x29`: Yellow gradual change
1. `0x2A`: Cyan gradual change
1. `0x2B`: Purple gradual change
1. `0x2C`: White gradual change
1. `0x2D`: Red, Green cross fade
1. `0x2E`: Red blue cross fade
1. `0x2F`: Green blue cross fade
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

Some operational modes take a speed parameter that controls how fast the colors are changed. `0x01` is the fastest, `0xFF` is the slowest.
