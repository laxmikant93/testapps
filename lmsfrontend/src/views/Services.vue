<!-- src/views/Services.vue -->
<template>
    <div class="services">
      <h1>Our Services</h1>
      <ul>
        <li>Web Development</li>
        <li>Mobile App Development</li>
        <li>Cloud Consulting</li>
        <li>Software Testing</li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Services',
  };
  </script>
  
  <style scoped>
  .services {
    padding: 20px;
    text-align: center;
  }
  .services ul {
    list-style: none;
    padding: 0;
  }
  .services ul li {
    padding: 10px 0;
  }
  </style>
  