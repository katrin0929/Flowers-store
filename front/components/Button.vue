<script>
export default {
  name: 'Button',
  props: {
    text: String
  },

  methods: {
    onClick() {
      this.$emit("submit")
    }
  }
}
</script>

<template>
  <button
    @click="onClick"
  >
    {{ text }}
  </button>
</template>

<style>
  button {
    background-color: #FF8F52;
    color: #FFFFFF;
    font-size: 20px;
    width: 390px;
    height: 57px;
    border-radius: 5px;
  }
</style>