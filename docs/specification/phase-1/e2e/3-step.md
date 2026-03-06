> Homework 25  Quiz questions CRUD by SA  Questions body validation  PUT -> "/sa/quiz/questions/:id/publish": should return error if passed body is incorrect; status 400;



    Tested endpoint: sa/quiz/questions/c3a69916-dc1e-4e3b-935a-700f5c9f3319/publish

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

> Homework 25  Access right for game flow  GET -> "/pair-game-quiz/pairs/my-current/answers": create new game by user1, connect by user2, try to add answer by user3. Should return error if current user is not inside active pair; status 403;  used additional methods: DELETE -> /testing/all-data, POST -> /sa/users, POST -> /auth/login, POST -> /sa/quiz/questions, PUT -> /sa/quiz/questions/:questionId/publish;



    Tested case params:
     endpoint: pair-game-quiz/pairs/my-current/answers,
     method: POST,
     headers: {"Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWU2OWU1OWMtYjQyZS00OGU0LTk0N2UtOGI3Y2E2NDIwNGZiIiwiZW1haWwiOiI3MDMxdXNlckBlbS5jb20iLCJ1c2VyX3R5cGUiOiJ1c2VyIiwiaXNzIjoib25saW5lLXF1aXotZ2FtZSIsImV4cCI6MTc3MjcxODMzNCwiaWF0IjoxNzcyNzE3NDM0fQ.CIBTu6k7qwjrFEoVs_60R5ki4vvRD_cmA9BwbjYQY1E"},
    body: {"answer":"correct answer"}

    Expected status: 403

    Received status: 404

    Description: try to update or delete the entity that was created by another user

      49 |       }
      50 |
    > 51 |       expect(status).toBeWith4xxStatus(expectedStatusCode, {
         |                      ^
      52 |         headers,
      53 |         endpoint,
      54 |         method,

      at performEndpointErrorChecker (src/tests/jest/back/testHelpers/performCheckers.ts:51:22)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-accessRight-describe.ts:239:7)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": add answers to first game, created by user1, connected by user2:
add correct answer by firstPlayer;
add incorrect answer by secondPlayer;
add correct answer by secondPlayer;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    AxiosError: request timeout of 6000ms exceeded

      46 |
      47 | export const addAnswer = async (answer: string, token: string) =>
    > 48 |   await axiosInstance.post<IAnswer>(
         |   ^
      49 |     PairQuizGameEndpoint.Answers,
      50 |     { answer },
      51 |     {

      at RedirectableRequest.handleRequestTimeout (node_modules/axios/lib/adapters/http.js:675:16)
      at Timeout.<anonymous> (node_modules/follow-redirects/index.js:221:12)
      at Axios.request (node_modules/axios/lib/core/Axios.js:45:41)
      at addAnswer (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:48:3)
      at performAddAnswerTestFlow (src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:130:72)
      at performAddAnswersTestFlow (src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:212:16)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:241:37)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create second game by user3, connect to the game by user4, then:
add correct answer by firstPlayer;
add incorrect answer by secondPlayer;
add correct answer by secondPlayer;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;

    request timeout of 6000ms exceeded

      25 |   }
      26 |
    > 27 |   throw new Error(error.message);
         |         ^
      28 | };
      29 |

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:27:9)
      at src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:33:91
      at createNewGameTestFlow (src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:33:5)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:189:7)

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

    request timeout of 6000ms exceeded

      25 |   }
      26 |
    > 27 |   throw new Error(error.message);
         |         ^
      28 | };
      29 |

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:27:9)
      at src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:85:26
      at getGameTestsFlow (src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:84:9)
      at performAddAnswersTestFlow (src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:220:7)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:241:37)

> Homework 25  Create, connect games, add answers  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create third game by user2, connect to the game by user1, then:
add correct answer by firstPlayer;
add incorrect answer by secondPlayer;
add correct answer by secondPlayer;
get active game and call "/pair-game-quiz/pairs/my-current by both users after each answer"; status 200;



    Expected: success response

    Received: Request failed with status code 403

    Config:
     url: pair-game-quiz/pairs/connection
     method: post
     response status: 403
     request body: {}
     response data: {"errorsMessages":[{"message":"You are already participating in an active game","field":""}]}

      22 |     };
      23 |
    > 24 |     expect(mappedError).printError(description);
         |                         ^
      25 |   }
      26 |
      27 |   throw new Error(error.message);

      at handleTestError (src/tests/jest/back/testHelpers/handleTestError.ts:24:25)
      at src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:33:91
      at createNewGameTestFlow (src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:33:5)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:189:7)

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

    expect(received).toStrictEqual(expected) // deep equality

    - Expected  -  4
    + Received  + 32

    @@ -8,10 +8,38 @@
          },
          "score": 0,
        },
        "id": Any<String>,
        "pairCreatedDate": StringMatching /\d{4}-[01]\d-[0-3]\dT[0-2]\d:[0-5]\d:[0-5]\d\.\d+([+-][0-2]\d:[0-5]\d|Z)/,
    -   "questions": null,
    -   "secondPlayerProgress": null,
    -   "startGameDate": null,
    -   "status": "PendingSecondPlayer",
    +   "questions": Array [
    +     Object {
    +       "body": "question body2 528502",
    +       "id": "479050d0-2c0b-4166-a66d-7729ab277fb5",
    +     },
    +     Object {
    +       "body": "question body3 529759",
    +       "id": "45ae1ca9-e90d-4bb0-b544-4904873467fc",
    +     },
    +     Object {
    +       "body": "question body4 531078",
    +       "id": "3541a0e5-ceb3-46bc-9b30-0fbf324e53de",
    +     },
    +     Object {
    +       "body": "question body0 525983",
    +       "id": "2622625c-2a13-44bd-8f0d-5ffa966d13ba",
    +     },
    +     Object {
    +       "body": "question body1 527257",
    +       "id": "e15e128d-b23a-4733-b6e9-c7105c791045",
    +     },
    +   ],
    +   "secondPlayerProgress": Object {
    +     "answers": Array [],
    +     "player": Object {
    +       "id": "7bde4658-7eb9-4683-a580-b7f09ae88435",
    +       "login": "1409lg",
    +     },
    +     "score": 0,
    +   },
    +   "startGameDate": "2026-03-05T16:32:50.23493+03:00",
    +   "status": "Active",
      }

      34 |
      35 |   expect(createNewGameResponseStatus).toBe(200);
    > 36 |   expect(createNewGameResponse).toStrictEqual(createNewGameSchema);
         |                                 ^
      37 |   expect(createNewGameResponse.firstPlayerProgress.player.id).toBe(userWhoCreateGame.userId);
      38 |   expect(createNewGameResponse.firstPlayerProgress.player.login).toBe(userWhoCreateGame.login);
      39 |

      at createNewGameTestFlow (src/tests/jest/back/describes/quiz/public/pairGameTestsFlow.ts:36:33)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:189:7)

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

      19 |   getSecondGame: (): IActivePairGame => {
      20 |     const secondGameData = expect.getState().secondGame;
    > 21 |     expect(secondGameData).not.toBeUndefined();
         |                                ^
      22 |
      23 |     return secondGameData;
      24 |   },

      at getExistingGameFromStateCallback (src/tests/jest/back/testHelpers/jestState/pairGameState.ts:21:32)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:239:28)

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

      31 |   getThirdGame: (): IActivePairGame => {
      32 |     const thirdGameData = expect.getState().thirdGame;
    > 33 |     expect(thirdGameData).not.toBeUndefined();
         |                               ^
      34 |
      35 |     return thirdGameData;
      36 |   },

      at getExistingGameFromStateCallback (src/tests/jest/back/testHelpers/jestState/pairGameState.ts:33:31)
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizeGame-createConnectGame-describe.ts:239:28)
