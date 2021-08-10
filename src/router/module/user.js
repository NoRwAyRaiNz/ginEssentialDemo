const userRoutes = [
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/register/register.vue'),
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/login.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true,
    },
    component: () => import('@/views/profile/Profile.vue'),

  },
];
export default userRoutes;
