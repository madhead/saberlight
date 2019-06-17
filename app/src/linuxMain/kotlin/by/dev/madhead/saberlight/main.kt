package by.dev.madhead.saberlight

import org.freedesktop.dbus.DBusBusType
import org.freedesktop.dbus.dbus_bus_get

fun main() {
    val bus = dbus_bus_get(DBusBusType.DBUS_BUS_SYSTEM, null)

    println(bus)

    println("Hello, world!")
}
