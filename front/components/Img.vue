<script>
export default {
  props: {
      alt: String,
      src: String,
      width: {
          type: String,
          default: '',
      },
      height: {
          type: String,
          default: '',
      },
  }, 
}
</script>

<template>
    <img
      :src
    />
</template>

