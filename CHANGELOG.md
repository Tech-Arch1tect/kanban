# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

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
