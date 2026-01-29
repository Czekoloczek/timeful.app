<template>
  <div class="tw-flex tw-flex-col tw-gap-2">
    <div class="tw-text-sm tw-font-medium tw-text-dark-green dark:tw-text-white">
      Theme
    </div>
    <v-menu offset-y>
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          class="tw-text-black dark:tw-text-white"
          outlined
          v-bind="attrs"
          v-on="on"
        >
          Theme: {{ themeLabel }}
        </v-btn>
      </template>
      <v-list dense class="dark:tw-bg-[#1b1e24]">
        <v-list-item @click="setTheme('system')">
          <v-list-item-title>System</v-list-item-title>
        </v-list-item>
        <v-list-item @click="setTheme('light')">
          <v-list-item-title>Light</v-list-item-title>
        </v-list-item>
        <v-list-item @click="setTheme('dark')">
          <v-list-item-title>Dark</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<script>
export default {
  name: "ThemeSelector",

  props: {
    themePreference: { type: String, required: true },
  },

  computed: {
    themeLabel() {
      if (this.themePreference === "dark") return "Dark"
      if (this.themePreference === "light") return "Light"
      return "System"
    },
  },

  methods: {
    setTheme(value) {
      this.$emit("update:themePreference", value)
    },
  },
}
</script>
