// src/router.js

import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import About from './views/About.vue';
import Services from './views/Services.vue';
import Contact from './views/Contact.vue';
import SignIn from './views/SignIn.vue';
import SignUp from './views/SignUp.vue';
import Dashboard from './views/Dashboard.vue'; // Protected route example

Vue.use(Router);

const routes = [
  // Public routes
  { path: '/', name: 'Home', component: Home },
  { path: '/about', name: 'About', component: About },
  { path: '/services', name: 'Services', component: Services },
  { path: '/contact', name: 'Contact', component: Contact },
  { path: '/signin', name: 'SignIn', component: SignIn },
  { path: '/signup', name: 'SignUp', component: SignUp },
  
  // Protected route example
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true },
  },
];

const router = new Router({
  routes,
  mode: 'history',
});

// Route guard to protect routes
router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  const token = localStorage.getItem('token');

  if (requiresAuth && !token) {
    next('/signin');
  } else {
    next();
  }
});

export default router;
