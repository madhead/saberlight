import org.gradle.api.internal.FeaturePreviews

enableFeaturePreview(FeaturePreviews.Feature.GRADLE_METADATA.name)

rootProject.name = "saberlight"

include(
        ":app"
)
