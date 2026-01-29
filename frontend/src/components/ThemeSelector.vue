<template>
  <div class="tw-flex tw-flex-col" :class="compact ? 'tw-gap-1' : 'tw-gap-2'">
    <div
      class="tw-font-medium tw-text-dark-green dark:tw-text-white"
      :class="compact ? 'tw-text-xs' : 'tw-text-sm'"
    >
      Theme
    </div>
    <v-menu offset-y>
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          class="tw-text-black dark:tw-text-white"
          :small="compact"
          :class="compact ? 'tw-text-xs tw-px-2' : ''"
          outlined
          v-bind="attrs"
          v-on="on"
        >
          Theme: {{ themeLabel }}
        </v-btn>
      </template>
      <v-list :dense="compact" class="dark:tw-bg-[#1b1e24]">
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
    compact: { type: Boolean, default: false },
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
