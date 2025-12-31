<template>
  <div>
    <h1>Test</h1>
    <div class="tw-my-4 tw-w-80">
      <AvailabilityTypeToggle v-model="availabilityType" />
    </div>
    <UpgradeDialog :value="true" />
  </div>
</template>

<script>
import UpgradeDialog from "@/components/pricing/UpgradeDialog"
import AvailabilityTypeToggle from "@/components/schedule_overlap/AvailabilityTypeToggle.vue"
import { availabilityTypes } from "@/constants"

export default {
  name: "Test",
  components: {
    UpgradeDialog,
    AvailabilityTypeToggle,
  },
  data() {
    return {
      availabilityType: availabilityTypes.AVAILABLE,
    }
  },
  methods: {
    createCheckoutSession() {
      post("/stripe/create-checkout-session").then((res) => {
        console.log(res)
      })
    },
  },
}
</script>
