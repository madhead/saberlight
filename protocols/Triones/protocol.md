# The droids we are looking for

<img src="../ZJ-MBL-RGBW (v3)/sample.png" alt="Subject appearance" width="100%">

This document describes GATT characteristics used to control Triones smart RGBW BLE bulbs (like on the picture above). Yes, they look the same as [Zengge bulbs](../ZJ-MBL-RGBW (v3)/protocol.md). My guess is that all cheap Chinese bulbs looks identically to each other, and only differ in GATT characteristics UUIDs.

All the text below is a result of my own experience of reverse-engineering its protocol. Use it at your own risk, **I am not responsible for your hardware.**

# Communication principles

TODO

# The protocol

## Status

TODO

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

TODO

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

Some operational modes take a speed parameter that controls how fast the colors are changed. `0x01` is the fastest, `0xFF` is the slowest.
