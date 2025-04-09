<script>

export default {
  props: {
    placeholder: String,
    modelValue: String
  },  

  methods: {
    updateModelValue(event) {
      this.$emit("update:modelValue", event.target.value)
    }
  }
}

</script>

<template>
    <input 
        @input="updateModelValue"
        :value="modelValue"
        :placeholder
    />
</template>

<style scoped>
    input {
        border-radius: 5px;
        border-width: 1px;
        border: 1px solid #D9D9D9;
        padding: 21px 0 21px 30px;   
    }
</style>