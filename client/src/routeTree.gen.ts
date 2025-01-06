/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as RegisterImport } from './routes/register'
import { Route as PasswordResetImport } from './routes/password-reset'
import { Route as LoginImport } from './routes/login'
import { Route as AboutImport } from './routes/about'
import { Route as IndexImport } from './routes/index'
import { Route as TaskIdImport } from './routes/task/$id'
import { Route as ProfileProfileImport } from './routes/profile/profile'
import { Route as Profile2faImport } from './routes/profile/2fa'
import { Route as BoardsSlugImport } from './routes/boards/$slug'
import { Route as AdminUsersImport } from './routes/admin/users'
import { Route as AdminBoardsImport } from './routes/admin/boards'
import { Route as BoardsAdministrateBoardIdImport } from './routes/boards/administrate/$boardId'

// Create/Update Routes

const RegisterRoute = RegisterImport.update({
  id: '/register',
  path: '/register',
  getParentRoute: () => rootRoute,
} as any)

const PasswordResetRoute = PasswordResetImport.update({
  id: '/password-reset',
  path: '/password-reset',
  getParentRoute: () => rootRoute,
} as any)

const LoginRoute = LoginImport.update({
  id: '/login',
  path: '/login',
  getParentRoute: () => rootRoute,
} as any)

const AboutRoute = AboutImport.update({
  id: '/about',
  path: '/about',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const TaskIdRoute = TaskIdImport.update({
  id: '/task/$id',
  path: '/task/$id',
  getParentRoute: () => rootRoute,
} as any)

const ProfileProfileRoute = ProfileProfileImport.update({
  id: '/profile/profile',
  path: '/profile/profile',
  getParentRoute: () => rootRoute,
} as any)

const Profile2faRoute = Profile2faImport.update({
  id: '/profile/2fa',
  path: '/profile/2fa',
  getParentRoute: () => rootRoute,
} as any)

const BoardsSlugRoute = BoardsSlugImport.update({
  id: '/boards/$slug',
  path: '/boards/$slug',
  getParentRoute: () => rootRoute,
} as any)

const AdminUsersRoute = AdminUsersImport.update({
  id: '/admin/users',
  path: '/admin/users',
  getParentRoute: () => rootRoute,
} as any)

const AdminBoardsRoute = AdminBoardsImport.update({
  id: '/admin/boards',
  path: '/admin/boards',
  getParentRoute: () => rootRoute,
} as any)

const BoardsAdministrateBoardIdRoute = BoardsAdministrateBoardIdImport.update({
  id: '/boards/administrate/$boardId',
  path: '/boards/administrate/$boardId',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutImport
      parentRoute: typeof rootRoute
    }
    '/login': {
      id: '/login'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof LoginImport
      parentRoute: typeof rootRoute
    }
    '/password-reset': {
      id: '/password-reset'
      path: '/password-reset'
      fullPath: '/password-reset'
      preLoaderRoute: typeof PasswordResetImport
      parentRoute: typeof rootRoute
    }
    '/register': {
      id: '/register'
      path: '/register'
      fullPath: '/register'
      preLoaderRoute: typeof RegisterImport
      parentRoute: typeof rootRoute
    }
    '/admin/boards': {
      id: '/admin/boards'
      path: '/admin/boards'
      fullPath: '/admin/boards'
      preLoaderRoute: typeof AdminBoardsImport
      parentRoute: typeof rootRoute
    }
    '/admin/users': {
      id: '/admin/users'
      path: '/admin/users'
      fullPath: '/admin/users'
      preLoaderRoute: typeof AdminUsersImport
      parentRoute: typeof rootRoute
    }
    '/boards/$slug': {
      id: '/boards/$slug'
      path: '/boards/$slug'
      fullPath: '/boards/$slug'
      preLoaderRoute: typeof BoardsSlugImport
      parentRoute: typeof rootRoute
    }
    '/profile/2fa': {
      id: '/profile/2fa'
      path: '/profile/2fa'
      fullPath: '/profile/2fa'
      preLoaderRoute: typeof Profile2faImport
      parentRoute: typeof rootRoute
    }
    '/profile/profile': {
      id: '/profile/profile'
      path: '/profile/profile'
      fullPath: '/profile/profile'
      preLoaderRoute: typeof ProfileProfileImport
      parentRoute: typeof rootRoute
    }
    '/task/$id': {
      id: '/task/$id'
      path: '/task/$id'
      fullPath: '/task/$id'
      preLoaderRoute: typeof TaskIdImport
      parentRoute: typeof rootRoute
    }
    '/boards/administrate/$boardId': {
      id: '/boards/administrate/$boardId'
      path: '/boards/administrate/$boardId'
      fullPath: '/boards/administrate/$boardId'
      preLoaderRoute: typeof BoardsAdministrateBoardIdImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexRoute
  '/about': typeof AboutRoute
  '/login': typeof LoginRoute
  '/password-reset': typeof PasswordResetRoute
  '/register': typeof RegisterRoute
  '/admin/boards': typeof AdminBoardsRoute
  '/admin/users': typeof AdminUsersRoute
  '/boards/$slug': typeof BoardsSlugRoute
  '/profile/2fa': typeof Profile2faRoute
  '/profile/profile': typeof ProfileProfileRoute
  '/task/$id': typeof TaskIdRoute
  '/boards/administrate/$boardId': typeof BoardsAdministrateBoardIdRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexRoute
  '/about': typeof AboutRoute
  '/login': typeof LoginRoute
  '/password-reset': typeof PasswordResetRoute
  '/register': typeof RegisterRoute
  '/admin/boards': typeof AdminBoardsRoute
  '/admin/users': typeof AdminUsersRoute
  '/boards/$slug': typeof BoardsSlugRoute
  '/profile/2fa': typeof Profile2faRoute
  '/profile/profile': typeof ProfileProfileRoute
  '/task/$id': typeof TaskIdRoute
  '/boards/administrate/$boardId': typeof BoardsAdministrateBoardIdRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexRoute
  '/about': typeof AboutRoute
  '/login': typeof LoginRoute
  '/password-reset': typeof PasswordResetRoute
  '/register': typeof RegisterRoute
  '/admin/boards': typeof AdminBoardsRoute
  '/admin/users': typeof AdminUsersRoute
  '/boards/$slug': typeof BoardsSlugRoute
  '/profile/2fa': typeof Profile2faRoute
  '/profile/profile': typeof ProfileProfileRoute
  '/task/$id': typeof TaskIdRoute
  '/boards/administrate/$boardId': typeof BoardsAdministrateBoardIdRoute
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
    | '/profile/profile'
    | '/task/$id'
    | '/boards/administrate/$boardId'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexRoute: typeof IndexRoute
  AboutRoute: typeof AboutRoute
  LoginRoute: typeof LoginRoute
  PasswordResetRoute: typeof PasswordResetRoute
  RegisterRoute: typeof RegisterRoute
  AdminBoardsRoute: typeof AdminBoardsRoute
  AdminUsersRoute: typeof AdminUsersRoute
  BoardsSlugRoute: typeof BoardsSlugRoute
  Profile2faRoute: typeof Profile2faRoute
  ProfileProfileRoute: typeof ProfileProfileRoute
  TaskIdRoute: typeof TaskIdRoute
  BoardsAdministrateBoardIdRoute: typeof BoardsAdministrateBoardIdRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexRoute: IndexRoute,
  AboutRoute: AboutRoute,
  LoginRoute: LoginRoute,
  PasswordResetRoute: PasswordResetRoute,
  RegisterRoute: RegisterRoute,
  AdminBoardsRoute: AdminBoardsRoute,
  AdminUsersRoute: AdminUsersRoute,
  BoardsSlugRoute: BoardsSlugRoute,
  Profile2faRoute: Profile2faRoute,
  ProfileProfileRoute: ProfileProfileRoute,
  TaskIdRoute: TaskIdRoute,
  BoardsAdministrateBoardIdRoute: BoardsAdministrateBoardIdRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* prettier-ignore-end */

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
        "/profile/profile",
        "/task/$id",
        "/boards/administrate/$boardId"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/about": {
      "filePath": "about.tsx"
    },
    "/login": {
      "filePath": "login.tsx"
    },
    "/password-reset": {
      "filePath": "password-reset.tsx"
    },
    "/register": {
      "filePath": "register.tsx"
    },
    "/admin/boards": {
      "filePath": "admin/boards.tsx"
    },
    "/admin/users": {
      "filePath": "admin/users.tsx"
    },
    "/boards/$slug": {
      "filePath": "boards/$slug.tsx"
    },
    "/profile/2fa": {
      "filePath": "profile/2fa.tsx"
    },
    "/profile/profile": {
      "filePath": "profile/profile.tsx"
    },
    "/task/$id": {
      "filePath": "task/$id.tsx"
    },
    "/boards/administrate/$boardId": {
      "filePath": "boards/administrate/$boardId.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
