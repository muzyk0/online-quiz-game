> Homework 25  Quiz questions CRUD by SA  Questions body validation  POST -> "/sa/quiz/questions": should return error if passed body is incorrect; status 400;



    Tested endpoint: sa/quiz/questions

    Tested method: POST

    Expected status: 400

    Received status: 400

    Passed body: {"body":"length_1001-f60ohVv5JGi6YrffYqmZNvpeS81tgDtWoVBTM0rTdlev2l4wzSq5CS54cCsRa7KIkcqw0w9PDWU2ohhSwNrFn7w4Ft6LMJ2dg2BMp1cSmzbanmfx6Fk04MKPmLVLL4hclHzdvJG8PZaWjoVARzz3NEAjNomZEf8JUcPC1YIOfZ9d3dQHyMmNuYw4gQXdKTPEJKscx0HDlhjMevSXPcO9LzD3s3XdAdK78bBBhxq1LHvyzPrBWAm4CcTqvjFc7S7ZwgpMUpyUv5jQqSecp0eC32nH0oOfso70XWwNm9JTvgTlUGwLvlGqqFN6dm9Q1RRjSkm2ug3bGs1pUj5DyQkLf7MQ8N2HZF7DBibN5HcIdcDO5FkAvh34k2qx1gzzKLlqE06wU8sBnNSvbDHe08oSw9bRALQZFADmEoccJ4XYjUgkEueIZUv0gUkDLd2Ur68HBbepUPfLWAwAEl1qqbCRhgUs7O0ajGnLFT040Lf8eZLEfLAxcvlcbZjzK8IIqn5gV96XkMGNCDul04Jj0UK2CBt0ul4UPqdEhG0afADNA7v1wU8ySbkuKlH5D92Z4zP9J7N9G664067iM7KNr1x8RJf0CeBKIXTTjf2iwE3bqetEVSgdK95XuJb5lnuIpo1yt7T9cVVNksmxmH4EUp9XxFwafjlr5CffaY1Ni1ybyKvnxQUOv6o9txu4FIcK9qrlsA8bgprU0z3774VYPTfcmrjPR6H6y6Dx777A1jLTy4e4ifpzYBa3gjtUygW2gbcwAXcbsDPOoUTZjemr3nQ7qBwtZCa84XHuXa7RfLPRRLzZ6OKOGlB0Vwe35pUkxK3z6FdIO1baFv0xAzcfpCwTeVzsAnugu0mOS8C0q9pHqevSDPx50eCYp67Z0hGcPtHDbodwy4iAqm06lxLfGCDJpOkQD8iZsDqIDbKv3WvwDk6sCelk8KCHY62vRHs5E0txhAezHnCrq0BJQK3wCDwRyPk7FZtrK","correctAnswers":["correct"]}

    Expected data: { errorsMessages: [{ message: Any<String>, field: "body" }] }

    Received data: {"errorsMessages":[{"message":"must be at most 500 characters long","field":"Body"}]}

      40 |     }
      41 |
    > 42 |     expect(response).toBeError(400, incorrectFields, requestBody, {
         |                      ^
      43 |       endpoint: url,
      44 |       method,
      45 |     });

      at performTestsWithIncorrectBody (src/tests/jest/back/testHelpers/performTestsWithIncorrectBody.ts:42:22)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/SA/questions-body-validation.ts:44:7)

> Homework 25  Quiz questions CRUD by SA  Questions body validation  PUT -> "/sa/quiz/questions/:id": should return error if passed body is incorrect; status 400;



    Tested endpoint: sa/quiz/questions/3f4e9ea6-f34d-43ee-8631-72677efc5b39

    Tested method: PUT

    Expected status: 400

    Received status: 400

    Passed body: {"body":"length_1001-f60ohVv5JGi6YrffYqmZNvpeS81tgDtWoVBTM0rTdlev2l4wzSq5CS54cCsRa7KIkcqw0w9PDWU2ohhSwNrFn7w4Ft6LMJ2dg2BMp1cSmzbanmfx6Fk04MKPmLVLL4hclHzdvJG8PZaWjoVARzz3NEAjNomZEf8JUcPC1YIOfZ9d3dQHyMmNuYw4gQXdKTPEJKscx0HDlhjMevSXPcO9LzD3s3XdAdK78bBBhxq1LHvyzPrBWAm4CcTqvjFc7S7ZwgpMUpyUv5jQqSecp0eC32nH0oOfso70XWwNm9JTvgTlUGwLvlGqqFN6dm9Q1RRjSkm2ug3bGs1pUj5DyQkLf7MQ8N2HZF7DBibN5HcIdcDO5FkAvh34k2qx1gzzKLlqE06wU8sBnNSvbDHe08oSw9bRALQZFADmEoccJ4XYjUgkEueIZUv0gUkDLd2Ur68HBbepUPfLWAwAEl1qqbCRhgUs7O0ajGnLFT040Lf8eZLEfLAxcvlcbZjzK8IIqn5gV96XkMGNCDul04Jj0UK2CBt0ul4UPqdEhG0afADNA7v1wU8ySbkuKlH5D92Z4zP9J7N9G664067iM7KNr1x8RJf0CeBKIXTTjf2iwE3bqetEVSgdK95XuJb5lnuIpo1yt7T9cVVNksmxmH4EUp9XxFwafjlr5CffaY1Ni1ybyKvnxQUOv6o9txu4FIcK9qrlsA8bgprU0z3774VYPTfcmrjPR6H6y6Dx777A1jLTy4e4ifpzYBa3gjtUygW2gbcwAXcbsDPOoUTZjemr3nQ7qBwtZCa84XHuXa7RfLPRRLzZ6OKOGlB0Vwe35pUkxK3z6FdIO1baFv0xAzcfpCwTeVzsAnugu0mOS8C0q9pHqevSDPx50eCYp67Z0hGcPtHDbodwy4iAqm06lxLfGCDJpOkQD8iZsDqIDbKv3WvwDk6sCelk8KCHY62vRHs5E0txhAezHnCrq0BJQK3wCDwRyPk7FZtrK","correctAnswers":["correct"]}

    Expected data: { errorsMessages: [{ message: Any<String>, field: "body" }] }

    Received data: {"errorsMessages":[{"message":"must be at most 500 characters long","field":"Body"}]}

      40 |     }
      41 |
    > 42 |     expect(response).toBeError(400, incorrectFields, requestBody, {
         |                      ^
      43 |       endpoint: url,
      44 |       method,
      45 |     });

      at performTestsWithIncorrectBody (src/tests/jest/back/testHelpers/performTestsWithIncorrectBody.ts:42:22)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/SA/questions-body-validation.ts:66:7)

> Homework 25  Quiz questions CRUD by SA  Questions body validation  PUT -> "/sa/quiz/questions/:id/publish": should return error if passed body is incorrect; status 400;



    Tested endpoint: sa/quiz/questions/48cb79d7-4dfb-44c4-b6bb-3e7fca19d4c8/publish

    Tested method: PUT

    Expected status: 400

    Received status: 400

    Passed body: {"published":"true"}

    Expected data: { errorsMessages: [{ message: Any<String>, field: "published" }] }

    Received data: {"errorsMessages":[{"message":"Invalid request body","field":""}]}

      40 |     }
      41 |
    > 42 |     expect(response).toBeError(400, incorrectFields, requestBody, {
         |                      ^
      43 |       endpoint: url,
      44 |       method,
      45 |     });

      at performTestsWithIncorrectBody (src/tests/jest/back/testHelpers/performTestsWithIncorrectBody.ts:42:22)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/SA/questions-body-validation.ts:88:7)

> Homework 25  Access right for game flow  GET -> "/pair-game-quiz/pairs/connection": create new game by user1, connect to game by user2, try to connect by user1, user2. Should return error if current user is already participating in active pair; status 403;  used additional methods: DELETE -> /testing/all-data, POST -> /sa/users, POST -> /auth/login, POST -> /sa/quiz/questions, PUT -> /sa/quiz/questions/:questionId/publish, POST -> /pair-game-quiz/pairs/connection;

    socket hang up

      25 |   }
      26 |
    > 27 |   throw new Error(error.message);
         |         ^
      28 | };
      29 |

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:27:9)
      at src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:97:41
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:93:27)

> Homework 25  Access right for game flow  GET -> "/pair-game-quiz/pairs/connection": create new game by user1, try to connect by user1. Should return error if current user is already participating in active pair; status 403;  used additional methods: DELETE -> /testing/all-data, POST -> /sa/users, POST -> /auth/login, POST -> /sa/quiz/questions, PUT -> /sa/quiz/questions/:questionId/publish, POST -> /pair-game-quiz/pairs/connection;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: testing/all-data
     method: delete
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:157:60
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:157:7)

> Homework 25  Access right for game flow  GET -> "/pair-game-quiz/pairs/my-current/answers": create new game by user1, connect by user2, try to add answer by user3. Should return error if current user is not inside active pair; status 403;  used additional methods: DELETE -> /testing/all-data, POST -> /sa/users, POST -> /auth/login, POST -> /sa/quiz/questions, PUT -> /sa/quiz/questions/:questionId/publish;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: testing/all-data
     method: delete
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:206:60
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:206:7)

> Homework 25  Access right for game flow  GET -> "/pair-game-quiz/pairs/my-current/answers": create new game by user1, try to add answer by user1. Should return error if current user is not inside active pair; status 403;  used additional methods: DELETE -> /testing/all-data, POST -> /sa/users, POST -> /auth/login, POST -> /sa/quiz/questions, PUT -> /sa/quiz/questions/:questionId/publish, POST -> /pair-game-quiz/pairs/connection;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: testing/all-data
     method: delete
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:260:60
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:260:7)

> Homework 25  Access right for game flow  GET -> "/pair-game-quiz/pairs/my-current/answers": create new game by user1, connect to game by user2, add 6 answers by user1. Should return error if current user has already answered to all questions; status 403;  used additional methods: DELETE -> /testing/all-data, POST -> /sa/users, POST -> /auth/login, POST -> /sa/quiz/questions, PUT -> /sa/quiz/questions/:questionId/publish, POST -> /pair-game-quiz/pairs/connection, POST -> /pair-game-quiz/pairs/my-current/answers;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: testing/all-data
     method: delete
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:312:60
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:312:7)

> Homework 25  Access right for game flow  GET -> "/pair-game-quiz/pairs/my-current": create new game by user1, connect to game by user2, add all answers by users. Should return error if no active pair for current user; status 404;  used additional methods: DELETE -> /testing/all-data, POST -> /sa/users, POST -> /auth/login, POST -> /pair-game-quiz/pairs/connection, POST -> /sa/quiz/questions, PUT -> /sa/quiz/questions/:questionId/publish, POST -> /pair-game-quiz/pairs/my-current/answers;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: testing/all-data
     method: delete
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:373:60
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:373:7)

> Homework 25  Exceptions for game flow  DELETE -> "/testing/all-data": should remove all data; status 204;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: testing/all-data
     method: delete
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/common.ts:19:79
      at Object.<anonymous> (src/tests/jest/back/describes/common.ts:19:26)

> Homework 25  Exceptions for game flow  POST -> "/sa/users": should create new user; status 201; content: created user;  used additional methods: GET => /sa/users;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: sa/users?pageSize=50
     method: get
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/usersApi/users-V2-describe.ts:121:24
      at Object.<anonymous> (src/tests/jest/back/describes/usersApi/users-V2-describe.ts:120:46)

> Homework 25  Exceptions for game flow  POST -> "/auth/login": should sign in user; status 200; content: JWT 'access' token, JWT 'refresh' token in cookie (http only, secure);

    expect(received).not.toBeUndefined()

    Received: undefined

      23 |   getUserCreds: (): IUserPayload => {
      24 |     const userCreds = expect.getState().newUserCreds;
    > 25 |     expect(userCreds).not.toBeUndefined();
         |                           ^
      26 |
      27 |     return userCreds;
      28 |   },

      at Object.getUserCreds (src/tests/jest/back/testHelpers/jestState/usersState.ts:25:27)
      at Object.<anonymous> (src/tests/jest/back/describes/refreshToken/refreshToken-describe.ts:61:40)

> Homework 25  Exceptions for game flow  POST -> "/sa/quiz/questions", PUT -> "/sa/quiz/questions/:questionId/publish": should create and publish several questions; status 201; content: created question;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: sa/quiz/questions
     method: post
     response status: 503
     request body: {"body":"question body0 492355","correctAnswers":["correct answer"]}
     response data: "<h1>no tunnel here :(</h1>"

      25 |   }
      26 |
    > 27 |   throw new Error(error.message);
         |         ^
      28 | };
      29 |

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:27:9)
      at src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:105:24
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:104:7)

> Homework 25  Exceptions for game flow  GET -> "/pair-game-quiz/pairs/my-current": should return error if there is no active pair for current user; status 404;

    expect(received).not.toBeUndefined()

    Received: undefined

      68 |   getAccessToken: (): string => {
      69 |     const accessToken = expect.getState().accessToken;
    > 70 |     expect(accessToken).not.toBeUndefined();
         |                             ^
      71 |
      72 |     return accessToken;
      73 |   },

      at Object.getAccessToken (src/tests/jest/back/testHelpers/jestState/usersState.ts:70:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGameExceptions-describe.ts:33:33)

> Homework 25  Exceptions for game flow  GET -> "/pair-game-quiz/pairs/my-current": should return error if such game does not exist; status 404;

    expect(received).not.toBeUndefined()

    Received: undefined

      68 |   getAccessToken: (): string => {
      69 |     const accessToken = expect.getState().accessToken;
    > 70 |     expect(accessToken).not.toBeUndefined();
         |                             ^
      71 |
      72 |     return accessToken;
      73 |   },

      at Object.getAccessToken (src/tests/jest/back/testHelpers/jestState/usersState.ts:70:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGameExceptions-describe.ts:58:33)

> Homework 25  Exceptions for game flow  GET -> "/pair-game-quiz/pairs/my-current", GET -> "pair-game-quiz/pairs/:id", POST -> "pair-game-quiz/pairs/connection", POST -> "pair-game-quiz/pairs/my-current/answers": should return error if auth credentials is incorrect; status 404;  used additional methods: POST -> /pair-game-quiz/pairs/connection;

    expect(received).not.toBeUndefined()

    Received: undefined

      68 |   getAccessToken: (): string => {
      69 |     const accessToken = expect.getState().accessToken;
    > 70 |     expect(accessToken).not.toBeUndefined();
         |                             ^
      71 |
      72 |     return accessToken;
      73 |   },

      at Object.getAccessToken (src/tests/jest/back/testHelpers/jestState/usersState.ts:70:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGameExceptions-describe.ts:84:33)

> Homework 25  Exceptions for game flow  GET -> "/pair-game-quiz/pairs": should return error if id has invalid format; status 400;

    expect(received).not.toBeUndefined()

    Received: undefined

      68 |   getAccessToken: (): string => {
      69 |     const accessToken = expect.getState().accessToken;
    > 70 |     expect(accessToken).not.toBeUndefined();
         |                             ^
      71 |
      72 |     return accessToken;
      73 |   },

      at Object.getAccessToken (src/tests/jest/back/testHelpers/jestState/usersState.ts:70:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGameExceptions-describe.ts:114:33)

> Homework 25  Create, connect games, add answers  DELETE -> "/testing/all-data": should remove all data; status 204;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: testing/all-data
     method: delete
     response status: 503
     request body: undefined
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/common.ts:19:79
      at Object.<anonymous> (src/tests/jest/back/describes/common.ts:19:26)

> Homework 25  Create, connect games, add answers  POST -> "/sa/users", "/auth/login": should create and login 6 users; status 201; content: created users;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: sa/users
     method: post
     response status: 503
     request body: {"login":"2969lg","password":"qwertwwy1","email":"2969user@em.com"}
     response data: "<h1>no tunnel here :(</h1>"

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/testHelpers/performTestsFlow/performTestsFlow.ts:240:40
      at createAndLoginSeveralUsers (src/tests/jest/back/testHelpers/performTestsFlow/performTestsFlow.ts:230:5)
      at Object.<anonymous> (src/tests/jest/back/describes/common.ts:39:27)

> Homework 25  Create, connect games, add answers  POST -> "/sa/quiz/questions", PUT -> "/sa/quiz/questions/:questionId/publish": should create and publish several questions; status 201; content: created question;



    Expected: success response

    Received: Request failed with status code 503

    Config:
     url: sa/quiz/questions
     method: post
     response status: 503
     request body: {"body":"question body0 493266","correctAnswers":["correct answer"]}
     response data: "<h1>no tunnel here :(</h1>"

      25 |   }
      26 |
    > 27 |   throw new Error(error.message);
         |         ^
      28 | };
      29 |

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:27:9)
      at src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:105:24
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:104:7)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/connection", GET -> "/pair-game-quiz/pairs/:id", GET -> "/pair-game-quiz/pairs/my-current": create new active game by user1, then get the game by user1, then call "/pair-game-quiz/pairs/my-current" by user1. Should return new created active game; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:118:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/connection", GET -> "/pair-game-quiz/pairs/:id", GET -> "/pair-game-quiz/pairs/my-current": connect to existing game by user2; then get the game by user1, user2; then call "/pair-game-quiz/pairs/my-current" by user1, user2. Should return started game; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:142:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": add answers to first game, created by user1, connected by user2:
add correct answer by firstPlayer;
add incorrect answer by secondPlayer;
add correct answer by secondPlayer;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:237:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create second game by user3, connect to the game by user4, then:
add correct answer by firstPlayer;
add incorrect answer by secondPlayer;
add correct answer by secondPlayer;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:187:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": add answers to first game, created by user1, connected by user2:
add correct answer by firstPlayer;
add correct answer by firstPlayer;
add correct answer by secondPlayer;
add correct answer by secondPlayer;
add incorrect answer by firstPlayer;
add correct answer by firstPlayer;
add correct answer by secondPlayer;
firstPlayer should win with 5 scores;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:237:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create third game by user2, connect to the game by user1, then:
add correct answer by firstPlayer;
add incorrect answer by secondPlayer;
add correct answer by secondPlayer;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:187:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create 4th game by user5, connect to the game by user6, then:
add correct answer by firstPlayer;
add incorrect answer by firstPlayer;
add correct answer by secondPlayer;
add incorrect answer by secondPlayer;
add incorrect answer by secondPlayer;
add incorrect answer by secondPlayer;
add incorrect answer by secondPlayer;
add correct answer by firstPlayer;
add incorrect answer by firstPlayer;
add incorrect answer by firstPlayer;
draw with 2 scores;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:187:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": add answers to second game, created by user3, connected by user4:
add incorrect answer by firstPlayer;
add incorrect answer by firstPlayer;
add correct answer by secondPlayer;
add correct answer by secondPlayer;
add incorrect answer by secondPlayer;
add correct answer by firstPlayer;
add incorrect answer by firstPlayer;
secondPlayer should win with 4 scores;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:237:45)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": add answers to third game, created by user2, connected by user1:
add correct answer by firstPlayer;
add correct answer by firstPlayer;
add correct answer by secondPlayer;
add correct answer by secondPlayer;
add incorrect answer by firstPlayer;
add correct answer by firstPlayer;
add correct answer by secondPlayer;
firstPlayer should win with 5 scores;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    expect(received).not.toBeUndefined()

    Received: undefined

       9 |   getUsers: (usersCount: number): ILoggedTestUser[] => {
      10 |     const loggedUsers: ILoggedTestUser[] = expect.getState().loggedUsers;
    > 11 |     expect(loggedUsers).not.toBeUndefined();
         |                             ^
      12 |     expect(loggedUsers.length).toBe(usersCount);
      13 |
      14 |     return loggedUsers;

      at Object.getUsers (src/tests/jest/back/testHelpers/jestState/usersState.ts:11:29)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:237:45)
