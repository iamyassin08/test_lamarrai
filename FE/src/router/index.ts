import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

import ProjectSlides from "@/components/projects/ProjectSlides.vue";



import "preline/preline";
import { type IStaticMethods } from "preline/preline";
import AppLayout from "../layouts/AppLayout.vue";


declare global {
  interface Window {
    HSStaticMethods: IStaticMethods;
  }
}
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      meta: {layout: AppLayout},
      component: HomeView,
    },

    
    
  ],
});

router.afterEach((to, from, failure) => {
  if (!failure) {
    setTimeout(() => {
      window.HSStaticMethods.autoInit();
    }, 100);
  }
});
export default router;