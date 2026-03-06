Suites (1)
1 failed
Tests (7)
3 passed
4 failed
/home/node/dist/auto-test-checker/dist/tests/jest/back/back-v3-28-homework-01/quiz-timing-for-game.spec.js
22.674s

Homework 28 > Finish game after waiting 10 sec.
DELETE -> "/testing/all-data": should remove all data; status 204;
passed
0.068s

Homework 28 > Finish game after waiting 10 sec.
POST -> "/sa/quiz/questions", PUT -> "/sa/quiz/questions/:questionId/publish": should create and publish several questions; status 201; content: created question;
passed
0.571s

Homework 28 > Finish game after waiting 10 sec.
POST -> "/sa/users", "/auth/login": should create and login 4 users; status 201; content: created users;
passed
7.298s

Homework 28 > Finish game after waiting 10 sec.
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game by user1, connect to game by user2. Add 5 correct answers by user1. Await 10 sec. Get game by user1. Should return finished game - status: "Finished", firstPlayerProgress.score: 6, secondPlayerProgress.score: 0, finishGameDate: not to be null; status 200;
failed
11.291s

Homework 28 > Finish game after waiting 10 sec.
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game by user1, connect to game by user2. Add 3 correct answers by user2. Add 5 correct answers by user1. Await 10 sec. Call "/pair-game-quiz/pairs/my-current" endpoint by user2. Should return status 404. Get game by user1. Should return finished game - status: "Finished", firstPlayerProgress.score: 6, secondPlayerProgress.score: 3, finishGameDate: not to be null; status 200;
failed
0.082s

Homework 28 > Finish game after waiting 10 sec.
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game by user1, connect to game by user2. Add 3 incorrect answers by user2. Add 3 correct answers by user1. Add 2 correct answers by user2. Call "/pair-game-quiz/pairs/my-current" endpoint by user2. Should return active game. Await 10 sec. Should return status 404. Get game by user2. Should return finished game - status: "Finished", firstPlayerProgress.score: 3, secondPlayerProgress.score: 3, finishGameDate: not to be null; status 200;
failed
0.051s

Homework 28 > Finish game after waiting 10 sec.
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game1 by user1, connect to game by user2. Add 3 incorrect answers by user2. Add 4 correct answers by user1. Create game2 by user3, connect to game by user4. Add 5 correct answers by user3. Add 2 correct answers by user4. Add 2 correct answers by user2. Await 10 sec. Get game1 by user2. Should return finished game - status: "Finished", firstPlayerProgress.score: 4, secondPlayerProgress.score: 3, finishGameDate: not to be null. Get game2 by user3. Should return finished game - status: "Finished", firstPlayerProgress.score: 6, secondPlayerProgress.score: 2, finishGameDate: not to be null. ; status 200;
failed
0.101s

> Homework 28  Finish game after waiting 10 sec.  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game by user1, connect to game by user2.
Add 5 correct answers by user1.
Await 10 sec.
Get game by user1.
Should return finished game -
status: "Finished",
firstPlayerProgress.score: 6,
secondPlayerProgress.score: 0,
finishGameDate: not to be null; status 200;

    expect(received).toBe(expected) // Object.is equality

    Expected: "Finished"
    Received: "Active"

      73 |       );
      74 |
    > 75 |       expect(activeGameAfterWaiting.status).toBe(GameStatus.Finished);
         |                                             ^
      76 |       expect(activeGameAfterWaiting.finishGameDate).not.toBeNull();
      77 |       expect(activeGameAfterWaiting.firstPlayerProgress.score).toBe(6); //plus extra point
      78 |       expect(activeGameAfterWaiting.secondPlayerProgress.score).toBe(0);

      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-timings-describe.ts:75:45)

> Homework 28  Finish game after waiting 10 sec.  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game by user1, connect to game by user2.
Add 3 correct answers by user2.
Add 5 correct answers by user1.
Await 10 sec.
Call "/pair-game-quiz/pairs/my-current" endpoint by user2.
Should return status 404.
Get game by user1.
Should return finished game -
status: "Finished",
firstPlayerProgress.score: 6,
secondPlayerProgress.score: 3,
finishGameDate: not to be null; status 200;



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
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-timings-describe.ts:105:7)

> Homework 28  Finish game after waiting 10 sec.  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game by user1, connect to game by user2.
Add 3 incorrect answers by user2.
Add 3 correct answers by user1.
Add 2 correct answers by user2.
Call "/pair-game-quiz/pairs/my-current" endpoint by user2.
Should return active game.
Await 10 sec.
Should return status 404.
Get game by user2.
Should return finished game -
status: "Finished",
firstPlayerProgress.score: 3,
secondPlayerProgress.score: 3,
finishGameDate: not to be null; status 200;



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
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-timings-describe.ts:172:7)

> Homework 28  Finish game after waiting 10 sec.  POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs": create game1 by user1, connect to game by user2.
Add 3 incorrect answers by user2.
Add 4 correct answers by user1.
Create game2 by user3, connect to game by user4.
Add 5 correct answers by user3.
Add 2 correct answers by user4.
Add 2 correct answers by user2.
Await 10 sec.
Get game1 by user2.
Should return finished game -
status: "Finished",
firstPlayerProgress.score: 4,
secondPlayerProgress.score: 3,
finishGameDate: not to be null.
Get game2 by user3.
Should return finished game -
status: "Finished",
firstPlayerProgress.score: 6,
secondPlayerProgress.score: 2,
finishGameDate: not to be null.
; status 200;



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
      at Object.<anonymous> (src/tests/jest/back/describes/quiz/public/pairQuizGame-timings-describe.ts:249:7)
