plugins {
    kotlin("multiplatform").version("1.3.31")
}

repositories {
    jcenter()
}

kotlin {
    //    linuxArm32Hfp("raspberrypi") {
    //        compilations {
    //            val main by getting {
    //                val dbus by cinterops.creating {
    //                    defFile(file("src/raspberrypiMain/cinterop/dbus/dbus.def"))
    //                    packageName("org.freedesktop.dbus")
    //                    includeDirs("src/raspberrypiMain/cinterop/dbus/headers")
    //                }
    //            }
    //        }
    //        binaries {
    //            executable("saberlight") {
    //                entryPoint = "by.dev.madhead.saberlight.main"
    //            }
    //        }
    //    }
    linuxX64("linux") {
        compilations {
            val main by getting {
                val dbus by cinterops.creating {
                    defFile(file("src/linuxMain/cinterop/dbus/dbus.def"))
                    packageName("org.freedesktop.dbus")
                    includeDirs(
                            "/usr/include/dbus-1.0/",
                            "/usr/lib/dbus-1.0/include/"
                    )
                }
            }
        }
        binaries {
            executable("saberlight") {
                entryPoint = "by.dev.madhead.saberlight.main"
                linkerOpts = mutableListOf(
                        "-L/usr/lib",
                        "-ldbus-1"
                )
            }
        }
    }
}
