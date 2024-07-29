export const base = '/'

export enum LAYOUTS {
    MAIN = 'layout-main',
    ADMIN = 'layout-admin'
}

export enum ROUTES {
    home = '',
    category = '/:category',
    product = '/product/:product',
    cart = '/cart',
    checkout = '/checkout',
    account = '/account',
    orders = '/account/orders',
    orderDetail = '/account/orders/:order'
}

export const routes: Record<
    string,
    {
        path: string
        name: string
        alias?: string
    }
> = {
    home: {
        name: 'home',
        path: ROUTES.home,
        alias: '/home'
    },
    category: {
        name: 'category',
        path: ROUTES.category
    },
    product: {
        name: 'product',
        path: ROUTES.product
    },
    cart: {
        name: 'cart',
        path: ROUTES.cart
    },
    checkout: {
        name: 'checkout',
        path: ROUTES.checkout
    },
    account: {
        name: 'account',
        path: ROUTES.account
    },
    orders: {
        name: 'orders',
        path: ROUTES.orders
    },
    orderDetail : {
        name: 'orderDetail',
        path: ROUTES.orderDetail
    }
}
