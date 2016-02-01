# The droids we are looking for

<img src="../ZJ-MBL-RGBW (v3)/sample.png" alt="Subject appearance" width="100%">

This document describes GATT characteristics used to control Triones smart RGBW BLE bulbs (like on the picture above). Yes, they look the same as [Zengge bulbs](../ZJ-MBL-RGBW (v3)). My guess is that all cheap Chinese bulbs looks identically to each other, and only differ in GATT characteristics UUIDs.

All the text below is a result of my own experience of reverse-engineering its protocol. Use it at your own risk, **I am not responsible for your hardware.**

# Communication principles

TODO

# The protocol

## Status

TODO

## Power

TODO

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
1. `payload[4]` _must_ be equal to magic constant `0x00` (TODO: It is used in "warm white" mode)
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

TODO

## Clock

TODO

# Magic constants

## Built-in modes

TODO

## Speed

TODO