import Vue from 'vue';
import VueRouter from 'vue-router';
// eslint-disable-next-line import/extensions
import Ping from '../components/Ping';

Vue.use(VueRouter);

const routes = [
  {
    path: '/ping',
    name: 'ping',
    component: Ping,
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
