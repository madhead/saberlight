plugins {
    kotlin("multiplatform").version("1.3.31")
}

repositories {
    jcenter()
}

kotlin {
    linuxArm32Hfp("raspberrypi") {
        binaries {
            executable("saberlight") {
                entryPoint = "by.dev.madhead.saberlight.main"
            }
        }
    }
}
