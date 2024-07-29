import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LayoutMain from '@/layouts/LayoutMain.vue'
import { base, LAYOUTS, routes } from './router'

const CategoryView = () => import('@/views/CategoryView.vue')
const ProductView = () => import('@/views/ProductView.vue')
const CartView = () => import('@/views/CartView.vue')
const CheckoutView = () => import('@/views/CheckoutView.vue')
const AccountView = () => import('@/views/AccountView.vue')
const AccountLayout = () => import('@/layouts/AccountLayout.vue')
const OrderView = () => import('@/views/OrderView.vue')
const OrderDetail = () => import('@/views/OrderDetailView.vue')

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: base,
            name: LAYOUTS.MAIN,
            component: LayoutMain,
            children: [
                {
                    path: routes.home.path,
                    name: routes.home.name,
                    alias: routes.home.alias,
                    component: HomeView
                },
                {
                    path: routes.category.path,
                    name: routes.category.name,
                    component: CategoryView
                },
                {
                    path: routes.product.path,
                    name: routes.product.name,
                    component: ProductView
                },
                {
                    path: routes.cart.path,
                    name: routes.cart.name,
                    component: CartView
                },
                {
                    path: routes.checkout.path,
                    name: routes.checkout.name,
                    component: CheckoutView
                },
                {
                    path: '',
                    component: AccountLayout,
                    children: [
                        {
                            path: routes.account.path,
                            name: routes.account.name,
                            component: AccountView
                        },
                        {
                            path: routes.orders.path,
                            name: routes.orders.name,
                            component: OrderView
                        }, {
                            path: routes.orderDetail.path,
                            name: routes.orderDetail.name,
                            component: OrderDetail
                        }
                    ]
                }
            ]
        },
        {
            path: '/about',
            name: 'about',
            component: () => import('@/views/AboutView.vue')
        }
    ]
})

export default router
