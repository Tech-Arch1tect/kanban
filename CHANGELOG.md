# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [0.3.12](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.11...v0.3.12) (2025-03-06)


### Features

* **taskColours:** Add 'TaskColourPicker' component ([f77e6ba](https://github.com/Tech-Arch1tect/kanban/commit/f77e6ba33991e4c3dbe525df2e3c3d4bc4313d61))
* **taskColours:** Add client hook for updating task colour ([cc3384b](https://github.com/Tech-Arch1tect/kanban/commit/cc3384b3578447de558a5aacd6e3581ad7ae401f))
* **taskColours:** Add colour property to task creation & add updateTaskColour route ([3f3a535](https://github.com/Tech-Arch1tect/kanban/commit/3f3a535c8d958d50260de840a3e4dcfbf38f8ccf))
* **taskColours:** add service layer for setting task.colour property ([2825ad6](https://github.com/Tech-Arch1tect/kanban/commit/2825ad665fafdcf0bd97b4aea5050114093e5e56))
* **taskColours:** Add task heading component for colour picking ([2271d4c](https://github.com/Tech-Arch1tect/kanban/commit/2271d4ca2345049bf9b8acb34231f88e837c40af))
* **taskColours:** Add utility for translating a string colour to tailwind classes ([eab250a](https://github.com/Tech-Arch1tect/kanban/commit/eab250a3722cd35607c4f9e008eabd3202fb9916))
* **taskColours:** Re-design Board Task Component - implementing a header which is coloured with task colour ([ec06fd4](https://github.com/Tech-Arch1tect/kanban/commit/ec06fd451358a1cd7190a0f04c1c9029665b7894))


### Bug Fixes

* **client:** Resolve "Cannot update a component (Lt) while rendering a different component (ViewFiles)" error when no file selected ([1825a4c](https://github.com/Tech-Arch1tect/kanban/commit/1825a4c3c12e94f2b5f68413f84d8ab3a2eb901e))
* **client:** useDownloadFile.ts fix some typescript possibly undefined errors ([6309b7f](https://github.com/Tech-Arch1tect/kanban/commit/6309b7fdc8311954422affbafd0733fb82ccf93c))


### Styles

* **columns:** Add small margin to the bottom of column titles ([6505f12](https://github.com/Tech-Arch1tect/kanban/commit/6505f122a9652b47a5225fbe4740c34751a9b4cd))
* **navbar:** Adjust navbar dark mode colour ([c3905c5](https://github.com/Tech-Arch1tect/kanban/commit/c3905c573376a7b1d9e24c6d44a305d144f4e8ed))


### Chores

* **client:** update ReactCompilerConfig target to 19 ([602b32e](https://github.com/Tech-Arch1tect/kanban/commit/602b32ebd4359f6c53a1431e10f2a84b599a3112))
* **screenshots:** Add screenshots to the repository ([1d46738](https://github.com/Tech-Arch1tect/kanban/commit/1d46738773c1ca8811a5fec9498fa1bae240c458))
* **taskColours:** Add tailwind bg classes to safelist ([a466121](https://github.com/Tech-Arch1tect/kanban/commit/a466121b39ce85b60c4469d8f4630f9fe5dfcf69))

### [0.3.11](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.10...v0.3.11) (2025-03-05)


### Chores

* **dependencies:** client: npm update ([52fe96c](https://github.com/Tech-Arch1tect/kanban/commit/52fe96cc8228eb432bd27ce1aaa35bda6bca243e))
* **Dependencies:** Switch from 'github.com/gin-gonic/contrib/sessions' to 'github.com/gin-contrib/sessions"' ([874b738](https://github.com/Tech-Arch1tect/kanban/commit/874b738eceb136598f2152e03889d6b444cad15e))
* **Dependencies:** switch from sessions.NewCookieStore -> cookie.NewStore ([7caeec7](https://github.com/Tech-Arch1tect/kanban/commit/7caeec729516c2da9c52bb0f7677bfcd3edc6bf7))
* **Dependencies:** Update go deps ([e76bacc](https://github.com/Tech-Arch1tect/kanban/commit/e76baccecf829a6bb05a0e4502b8be97c44d04cd))
* **react:** Upgrade client to react 19.x ([e025c65](https://github.com/Tech-Arch1tect/kanban/commit/e025c6551292dded60418ce0abb0da4783f389a7))


### Code Refactoring

* **Images:** Client: Cache images for 10 mins ([a11ae67](https://github.com/Tech-Arch1tect/kanban/commit/a11ae67d8321c1f9b962c070ecdaa23bd36318fd))
* **logging:** Move Gin over to Zap for logging ([b8858f2](https://github.com/Tech-Arch1tect/kanban/commit/b8858f2c49f6e5e1651faa6bfa1b38f572050700))
* **logging:** Move most "log" usage over to "zap" logger ([c9d6b8e](https://github.com/Tech-Arch1tect/kanban/commit/c9d6b8e6ed9c264478e314b1addc1a48b565fb52))

### [0.3.10](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.9...v0.3.10) (2025-03-04)


### Features

* **Mentions:** Add 'mentioned' event friendly name for client side ([f32d779](https://github.com/Tech-Arch1tect/kanban/commit/f32d77945382238472e6be7282d50607cb676ffb))
* **Mentions:** Add 'mentioned' event to seeded events ([25f6bd2](https://github.com/Tech-Arch1tect/kanban/commit/25f6bd2f17371186041809bbd179de4390a95f01))
* **Mentions:** Add notification subscriber for mention event ([8dcdcf6](https://github.com/Tech-Arch1tect/kanban/commit/8dcdcf6799e3c40e353054f767dc3240b26e75be))
* **Mentions:** Create event bus for task or comment changes + subscribe to comment & task events to detect and publish mention events ([c7ecfb9](https://github.com/Tech-Arch1tect/kanban/commit/c7ecfb969509309593fe464ecd7e85395978fc79))

### [0.3.9](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.8...v0.3.9) (2025-03-03)


### Bug Fixes

* **CommentItem:** Fix erroneous scrolling on comment hovering:  Move Reaction picker to the top of the comment ([51161c1](https://github.com/Tech-Arch1tect/kanban/commit/51161c181748f50b63403d724d28656184da230f))


### Chores

* **cleanup:** server: Correct NotificationEventRepository file name ([a8e31c1](https://github.com/Tech-Arch1tect/kanban/commit/a8e31c17b2d18bedc3832468e96ed4b52790a24f))


### Code Refactoring

* **database:** server: move Migrate() to the generic repository definition ([9403d76](https://github.com/Tech-Arch1tect/kanban/commit/9403d76e562e7db47cbe10622687e257ebcbf6a0))


### Styles

* **TaskActivity:** Compact the layout of task activities ([9da3713](https://github.com/Tech-Arch1tect/kanban/commit/9da37136be4bae7a46c2a86affb4edce0b1af69c))
* **TaskActivity:** Improve styling of data changes ([1c791d0](https://github.com/Tech-Arch1tect/kanban/commit/1c791d038cbd797270d06c642094026cfb8d8296))
* **TaskActivity:** Move pagination to the top of the task activities component ([e7311a4](https://github.com/Tech-Arch1tect/kanban/commit/e7311a4ee7cd82d3b6bb2df3ca0639ed9d615952))
* **TaskActivity:** remove padding to bring title in-line with other titles ([1b58795](https://github.com/Tech-Arch1tect/kanban/commit/1b58795b6323383f7bd50abc2b2cfc3281feab0c))
* **TaskActivity:** Render markdown for change data ([bc75f96](https://github.com/Tech-Arch1tect/kanban/commit/bc75f96554a9dc86f1776f5eb277a14e91984ebd))

### [0.3.8](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.7...v0.3.8) (2025-03-02)


### Features

* **TaskActivity:** Add functionality to Task service to retrieve paginated Task activities ([26ce6a8](https://github.com/Tech-Arch1tect/kanban/commit/26ce6a80273243b3a9ca739ad0e46e115b2a713b))
* **TaskActivity:** client: Add component for displaying Task Activities ([bd248f0](https://github.com/Tech-Arch1tect/kanban/commit/bd248f003786d4338a4bed53d2e493bd417eec3b))
* **TaskActivity:** client: Add hook to retrieve Task Activities by task ID (with pagination) ([2c48a39](https://github.com/Tech-Arch1tect/kanban/commit/2c48a392cbbc18d87bbc2e0bfdfa29facf3b3bb8))
* **TaskActivity:** server: Add GetTaskActivities route ([05a2027](https://github.com/Tech-Arch1tect/kanban/commit/05a2027fb650e2101a34dcd86a98d56a971ad993))
* **TaskActivity:** server: Add TaskActivity Model and DB repository ([04d7c1e](https://github.com/Tech-Arch1tect/kanban/commit/04d7c1ed09d0b5bbc2969da6ca80c8121ae440f5))
* **TaskActivity:** server: Implement TaskActivityService ([b33caa4](https://github.com/Tech-Arch1tect/kanban/commit/b33caa4d701534e125785809909a14a2bba96249))


### Bug Fixes

* **events:** server: correctly set new/old data in externallink.created event ([675e502](https://github.com/Tech-Arch1tect/kanban/commit/675e502ac5db9d79798344643e86e91a348eb9fa))
* **TaskActivity:** client: Correct task activity background colour ([6a6f1a2](https://github.com/Tech-Arch1tect/kanban/commit/6a6f1a26d6bb016cba46f2c173f1027cc064170d))
* **TaskActivity:** Delete task activity when task is deleted ([2e0156a](https://github.com/Tech-Arch1tect/kanban/commit/2e0156af392bb9932418719df34e414f4f94747c))


### Performance Improvements

* **react:** client: Switch to lazy loading for better production bundle splitting ([62af830](https://github.com/Tech-Arch1tect/kanban/commit/62af830afaec28643d94b1b769e17a8c1bdd1110))


### Code Refactoring

* **PaginatedSearch:** Add a generic PaginatedSearch method for db repositories and remove old implementation on users repo ([ac34dec](https://github.com/Tech-Arch1tect/kanban/commit/ac34dec44db730e635a2101413527a253de54cbb))

### [0.3.7](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.6...v0.3.7) (2025-03-01)

### [0.3.6](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.5...v0.3.6) (2025-03-01)


### Bug Fixes

* **ci:** docker build and push: Attempt to determine the tag from the checked out commit ([7ef2b1e](https://github.com/Tech-Arch1tect/kanban/commit/7ef2b1e328460761f7336ae565c2025141a3e4d6))

### [0.3.5](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.4...v0.3.5) (2025-03-01)


### Bug Fixes

* **CI:** Change docker-build workflow to be triggered by the release workflow ([022d9ce](https://github.com/Tech-Arch1tect/kanban/commit/022d9ce1aee9a5969257f9c08fa7d47b55ea6c5d))
* **client:** Change default page title to "Kanban" from "Tanstack Router" ([ca16b9c](https://github.com/Tech-Arch1tect/kanban/commit/ca16b9c1b9507672ac78d220e45e1837e75f2ebd))

### [0.3.4](https://github.com/Tech-Arch1tect/kanban/compare/v0.3.3...v0.3.4) (2025-02-28)

### 0.3.3 (2025-02-27)


### Code Refactoring

* **both:** Add configurations for standard-version for easier changelog management and releases ([1f02a04](https://github.com/Tech-Arch1tect/kanban/commit/1f02a043a1586f0c66de337973418b1da7af1821))


### Chores

* **ci:** Add manually triggered CI job to create a release ([3821bba](https://github.com/Tech-Arch1tect/kanban/commit/3821bba1245f468d1e751cd11e93c657e57032eb))
* **cleanup:** Remove changelog which will be re-generated upon next release ([5eae426](https://github.com/Tech-Arch1tect/kanban/commit/5eae426d09c3bdaaf03a57c31a6c6f94f091777f))
* **release:** 0.3.3 ([4632b5d](https://github.com/Tech-Arch1tect/kanban/commit/4632b5d0bb910fb206c404ea856c5499d464a23e))
* **server:** Update all event bus subscribers to accept a change instead of a type instance ([6a901f6](https://github.com/Tech-Arch1tect/kanban/commit/6a901f6a7226d559488eae05516cae48e60480b2))
* **server:** update comment service and controller to accommodate eventbus / notification subscriber changes ([3f664fe](https://github.com/Tech-Arch1tect/kanban/commit/3f664fe222f8c751bd5c5f249b18eb1324f75415))
* **server:** Update Generic Event Bus to take old and new data ([a0fb749](https://github.com/Tech-Arch1tect/kanban/commit/a0fb7492c351c3f8614f9f5de4a2eb768fbe5997))
* **server:** update task service and controller to accommodate eventbus / notification subscriber changes ([273280d](https://github.com/Tech-Arch1tect/kanban/commit/273280d9eee3571826d84a56a7ef124c3966649d))
