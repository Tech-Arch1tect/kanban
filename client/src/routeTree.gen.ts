/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'

// Create Virtual Routes

const RegisterLazyImport = createFileRoute('/register')()
const PasswordResetLazyImport = createFileRoute('/password-reset')()
const LoginLazyImport = createFileRoute('/login')()
const AboutLazyImport = createFileRoute('/about')()
const IndexLazyImport = createFileRoute('/')()
const TaskIdLazyImport = createFileRoute('/task/$id')()
const ProfileProfileLazyImport = createFileRoute('/profile/profile')()
const ProfileNotificationsLazyImport = createFileRoute(
  '/profile/notifications',
)()
const Profile2faLazyImport = createFileRoute('/profile/2fa')()
const BoardsSlugLazyImport = createFileRoute('/boards/$slug')()
const AdminUsersLazyImport = createFileRoute('/admin/users')()
const AdminBoardsLazyImport = createFileRoute('/admin/boards')()
const BoardsAdministrateBoardIdLazyImport = createFileRoute(
  '/boards/administrate/$boardId',
)()

// Create/Update Routes

const RegisterLazyRoute = RegisterLazyImport.update({
  id: '/register',
  path: '/register',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/register.lazy').then((d) => d.Route))

const PasswordResetLazyRoute = PasswordResetLazyImport.update({
  id: '/password-reset',
  path: '/password-reset',
  getParentRoute: () => rootRoute,
} as any).lazy(() =>
  import('./routes/password-reset.lazy').then((d) => d.Route),
)

const LoginLazyRoute = LoginLazyImport.update({
  id: '/login',
  path: '/login',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/login.lazy').then((d) => d.Route))

const AboutLazyRoute = AboutLazyImport.update({
  id: '/about',
  path: '/about',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/about.lazy').then((d) => d.Route))

const IndexLazyRoute = IndexLazyImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/index.lazy').then((d) => d.Route))

const TaskIdLazyRoute = TaskIdLazyImport.update({
  id: '/task/$id',
  path: '/task/$id',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/task/$id.lazy').then((d) => d.Route))

const ProfileProfileLazyRoute = ProfileProfileLazyImport.update({
  id: '/profile/profile',
  path: '/profile/profile',
  getParentRoute: () => rootRoute,
} as any).lazy(() =>
  import('./routes/profile/profile.lazy').then((d) => d.Route),
)

const ProfileNotificationsLazyRoute = ProfileNotificationsLazyImport.update({
  id: '/profile/notifications',
  path: '/profile/notifications',
  getParentRoute: () => rootRoute,
} as any).lazy(() =>
  import('./routes/profile/notifications.lazy').then((d) => d.Route),
)

const Profile2faLazyRoute = Profile2faLazyImport.update({
  id: '/profile/2fa',
  path: '/profile/2fa',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/profile/2fa.lazy').then((d) => d.Route))

const BoardsSlugLazyRoute = BoardsSlugLazyImport.update({
  id: '/boards/$slug',
  path: '/boards/$slug',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/boards/$slug.lazy').then((d) => d.Route))

const AdminUsersLazyRoute = AdminUsersLazyImport.update({
  id: '/admin/users',
  path: '/admin/users',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/admin/users.lazy').then((d) => d.Route))

const AdminBoardsLazyRoute = AdminBoardsLazyImport.update({
  id: '/admin/boards',
  path: '/admin/boards',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/admin/boards.lazy').then((d) => d.Route))

const BoardsAdministrateBoardIdLazyRoute =
  BoardsAdministrateBoardIdLazyImport.update({
    id: '/boards/administrate/$boardId',
    path: '/boards/administrate/$boardId',
    getParentRoute: () => rootRoute,
  } as any).lazy(() =>
    import('./routes/boards/administrate/$boardId.lazy').then((d) => d.Route),
  )

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexLazyImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutLazyImport
      parentRoute: typeof rootRoute
    }
    '/login': {
      id: '/login'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof LoginLazyImport
      parentRoute: typeof rootRoute
    }
    '/password-reset': {
      id: '/password-reset'
      path: '/password-reset'
      fullPath: '/password-reset'
      preLoaderRoute: typeof PasswordResetLazyImport
      parentRoute: typeof rootRoute
    }
    '/register': {
      id: '/register'
      path: '/register'
      fullPath: '/register'
      preLoaderRoute: typeof RegisterLazyImport
      parentRoute: typeof rootRoute
    }
    '/admin/boards': {
      id: '/admin/boards'
      path: '/admin/boards'
      fullPath: '/admin/boards'
      preLoaderRoute: typeof AdminBoardsLazyImport
      parentRoute: typeof rootRoute
    }
    '/admin/users': {
      id: '/admin/users'
      path: '/admin/users'
      fullPath: '/admin/users'
      preLoaderRoute: typeof AdminUsersLazyImport
      parentRoute: typeof rootRoute
    }
    '/boards/$slug': {
      id: '/boards/$slug'
      path: '/boards/$slug'
      fullPath: '/boards/$slug'
      preLoaderRoute: typeof BoardsSlugLazyImport
      parentRoute: typeof rootRoute
    }
    '/profile/2fa': {
      id: '/profile/2fa'
      path: '/profile/2fa'
      fullPath: '/profile/2fa'
      preLoaderRoute: typeof Profile2faLazyImport
      parentRoute: typeof rootRoute
    }
    '/profile/notifications': {
      id: '/profile/notifications'
      path: '/profile/notifications'
      fullPath: '/profile/notifications'
      preLoaderRoute: typeof ProfileNotificationsLazyImport
      parentRoute: typeof rootRoute
    }
    '/profile/profile': {
      id: '/profile/profile'
      path: '/profile/profile'
      fullPath: '/profile/profile'
      preLoaderRoute: typeof ProfileProfileLazyImport
      parentRoute: typeof rootRoute
    }
    '/task/$id': {
      id: '/task/$id'
      path: '/task/$id'
      fullPath: '/task/$id'
      preLoaderRoute: typeof TaskIdLazyImport
      parentRoute: typeof rootRoute
    }
    '/boards/administrate/$boardId': {
      id: '/boards/administrate/$boardId'
      path: '/boards/administrate/$boardId'
      fullPath: '/boards/administrate/$boardId'
      preLoaderRoute: typeof BoardsAdministrateBoardIdLazyImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexLazyRoute
  '/about': typeof AboutLazyRoute
  '/login': typeof LoginLazyRoute
  '/password-reset': typeof PasswordResetLazyRoute
  '/register': typeof RegisterLazyRoute
  '/admin/boards': typeof AdminBoardsLazyRoute
  '/admin/users': typeof AdminUsersLazyRoute
  '/boards/$slug': typeof BoardsSlugLazyRoute
  '/profile/2fa': typeof Profile2faLazyRoute
  '/profile/notifications': typeof ProfileNotificationsLazyRoute
  '/profile/profile': typeof ProfileProfileLazyRoute
  '/task/$id': typeof TaskIdLazyRoute
  '/boards/administrate/$boardId': typeof BoardsAdministrateBoardIdLazyRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexLazyRoute
  '/about': typeof AboutLazyRoute
  '/login': typeof LoginLazyRoute
  '/password-reset': typeof PasswordResetLazyRoute
  '/register': typeof RegisterLazyRoute
  '/admin/boards': typeof AdminBoardsLazyRoute
  '/admin/users': typeof AdminUsersLazyRoute
  '/boards/$slug': typeof BoardsSlugLazyRoute
  '/profile/2fa': typeof Profile2faLazyRoute
  '/profile/notifications': typeof ProfileNotificationsLazyRoute
  '/profile/profile': typeof ProfileProfileLazyRoute
  '/task/$id': typeof TaskIdLazyRoute
  '/boards/administrate/$boardId': typeof BoardsAdministrateBoardIdLazyRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexLazyRoute
  '/about': typeof AboutLazyRoute
  '/login': typeof LoginLazyRoute
  '/password-reset': typeof PasswordResetLazyRoute
  '/register': typeof RegisterLazyRoute
  '/admin/boards': typeof AdminBoardsLazyRoute
  '/admin/users': typeof AdminUsersLazyRoute
  '/boards/$slug': typeof BoardsSlugLazyRoute
  '/profile/2fa': typeof Profile2faLazyRoute
  '/profile/notifications': typeof ProfileNotificationsLazyRoute
  '/profile/profile': typeof ProfileProfileLazyRoute
  '/task/$id': typeof TaskIdLazyRoute
  '/boards/administrate/$boardId': typeof BoardsAdministrateBoardIdLazyRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths:
    | '/'
    | '/about'
    | '/login'
    | '/password-reset'
    | '/register'
    | '/admin/boards'
    | '/admin/users'
    | '/boards/$slug'
    | '/profile/2fa'
    | '/profile/notifications'
    | '/profile/profile'
    | '/task/$id'
    | '/boards/administrate/$boardId'
  fileRoutesByTo: FileRoutesByTo
  to:
    | '/'
    | '/about'
    | '/login'
    | '/password-reset'
    | '/register'
    | '/admin/boards'
    | '/admin/users'
    | '/boards/$slug'
    | '/profile/2fa'
    | '/profile/notifications'
    | '/profile/profile'
    | '/task/$id'
    | '/boards/administrate/$boardId'
  id:
    | '__root__'
    | '/'
    | '/about'
    | '/login'
    | '/password-reset'
    | '/register'
    | '/admin/boards'
    | '/admin/users'
    | '/boards/$slug'
    | '/profile/2fa'
    | '/profile/notifications'
    | '/profile/profile'
    | '/task/$id'
    | '/boards/administrate/$boardId'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexLazyRoute: typeof IndexLazyRoute
  AboutLazyRoute: typeof AboutLazyRoute
  LoginLazyRoute: typeof LoginLazyRoute
  PasswordResetLazyRoute: typeof PasswordResetLazyRoute
  RegisterLazyRoute: typeof RegisterLazyRoute
  AdminBoardsLazyRoute: typeof AdminBoardsLazyRoute
  AdminUsersLazyRoute: typeof AdminUsersLazyRoute
  BoardsSlugLazyRoute: typeof BoardsSlugLazyRoute
  Profile2faLazyRoute: typeof Profile2faLazyRoute
  ProfileNotificationsLazyRoute: typeof ProfileNotificationsLazyRoute
  ProfileProfileLazyRoute: typeof ProfileProfileLazyRoute
  TaskIdLazyRoute: typeof TaskIdLazyRoute
  BoardsAdministrateBoardIdLazyRoute: typeof BoardsAdministrateBoardIdLazyRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexLazyRoute: IndexLazyRoute,
  AboutLazyRoute: AboutLazyRoute,
  LoginLazyRoute: LoginLazyRoute,
  PasswordResetLazyRoute: PasswordResetLazyRoute,
  RegisterLazyRoute: RegisterLazyRoute,
  AdminBoardsLazyRoute: AdminBoardsLazyRoute,
  AdminUsersLazyRoute: AdminUsersLazyRoute,
  BoardsSlugLazyRoute: BoardsSlugLazyRoute,
  Profile2faLazyRoute: Profile2faLazyRoute,
  ProfileNotificationsLazyRoute: ProfileNotificationsLazyRoute,
  ProfileProfileLazyRoute: ProfileProfileLazyRoute,
  TaskIdLazyRoute: TaskIdLazyRoute,
  BoardsAdministrateBoardIdLazyRoute: BoardsAdministrateBoardIdLazyRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/about",
        "/login",
        "/password-reset",
        "/register",
        "/admin/boards",
        "/admin/users",
        "/boards/$slug",
        "/profile/2fa",
        "/profile/notifications",
        "/profile/profile",
        "/task/$id",
        "/boards/administrate/$boardId"
      ]
    },
    "/": {
      "filePath": "index.lazy.tsx"
    },
    "/about": {
      "filePath": "about.lazy.tsx"
    },
    "/login": {
      "filePath": "login.lazy.tsx"
    },
    "/password-reset": {
      "filePath": "password-reset.lazy.tsx"
    },
    "/register": {
      "filePath": "register.lazy.tsx"
    },
    "/admin/boards": {
      "filePath": "admin/boards.lazy.tsx"
    },
    "/admin/users": {
      "filePath": "admin/users.lazy.tsx"
    },
    "/boards/$slug": {
      "filePath": "boards/$slug.lazy.tsx"
    },
    "/profile/2fa": {
      "filePath": "profile/2fa.lazy.tsx"
    },
    "/profile/notifications": {
      "filePath": "profile/notifications.lazy.tsx"
    },
    "/profile/profile": {
      "filePath": "profile/profile.lazy.tsx"
    },
    "/task/$id": {
      "filePath": "task/$id.lazy.tsx"
    },
    "/boards/administrate/$boardId": {
      "filePath": "boards/administrate/$boardId.lazy.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
