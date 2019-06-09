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
                // entryPoint = "by.dev.madhead.saberlight.main"
                entryPoint = "io.ktor.server.cio.EngineMain"
            }
        }
    }

    sourceSets {
        val raspberrypiMain by getting {
            dependencies {
                implementation("io.ktor:ktor-server-cio:1.2.1")
            }
        }
    }
}
