Suites (1)
1 passed
Tests (16)
16 passed
/home/node/dist/auto-test-checker/dist/tests/jest/back/back-v3-26-homework-01/quiz-my-games-my-statistic-functionality.spec.js
31.639s

Homework 26 > My games
DELETE -> "/testing/all-data": should remove all data; status 204;
passed
0.173s

Homework 26 > My games
POST -> "/sa/quiz/questions", PUT -> "/sa/quiz/questions/:questionId/publish": should create and publish several questions; status 201; content: created question;
passed
1.307s

Homework 26 > My games
POST -> "/sa/users", "/auth/login": should create and login 2 users; status 201; content: created users;
passed
2.811s

Homework 26 > My games
GET -> "/pair-game-quiz/pairs/my": create, connect to game 3 times and add all answers sequantly. Then create, connect 4th game (not finished). Then get 'myGames' sorted by 'status'; status 200; content: list of current user games (finished and current); used additional methods: POST -> /pair-game-quiz/pairs/connection, POST -> /pair-game-quiz/pairs/my-current/answers;
passed
5.069s

Homework 26 > My statistic
DELETE -> "/testing/all-data": should remove all data; status 204;
passed
0.161s

Homework 26 > My statistic
POST -> "/sa/quiz/questions", PUT -> "/sa/quiz/questions/:questionId/publish": should create and publish several questions; status 201; content: created question;
passed
1.213s

Homework 26 > My statistic
POST -> "/sa/users", "/auth/login": should create and login 4 users; status 201; content: created users;
passed
5.607s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user2, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.524s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user2, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.524s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user2, connect to the game by user1, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.546s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user2, connect to the game by user1, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.528s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user3, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; draw with 2 scores; ; status 200;
passed
1.461s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user1, connect to the game by user4, then: add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; secondPlayer should win, scores: 5 - 0; no one got an extra point; ; status 200;
passed
1.832s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user4, connect to the game by user1, then: add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add correct answer by secondPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; secondPlayer should win, scores: 5 - 0; no one got an extra point; ; status 200;
passed
1.51s

Homework 26 > My statistic
POST -> "/pair-game-quiz/pairs/my-current/answers", GET -> "/pair-game-quiz/pairs", GET -> "/pair-game-quiz/pairs/my-current": create game by user2, connect to the game by user4, then: add correct answer by firstPlayer; add incorrect answer by firstPlayer; add correct answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add incorrect answer by secondPlayer; add correct answer by firstPlayer; add correct answer by firstPlayer; add incorrect answer by firstPlayer; firstPlayer should win, scores: 3 - 2; ; status 200;
passed
1.488s

Homework 26 > My statistic
GET -> "/pair-game-quiz/users/my-statistic": should return status 200; content: current user's games statistic;
passed
0.125s

⇨ http server started on [::]:8080
{"ip":"62.133.61.254","latency":"25.417µs","method":"GET","request_id":"RMQBTJiktZNtgsRdjEduEAZkpNgxqpXn","status":404,"time":"2026-03-05T23:17:48+03:00","uri":"/","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.3 Safari/605.1.15"}
{"time":"2026-03-05T23:17:48.598473+03:00","level":"ERROR","prefix":"echo","file":"middleware.go","line":"50","message":"HTTP error: 404 - / - Not Found"}
{"ip":"62.133.61.254","latency":"9.042µs","method":"GET","request_id":"DgmwCCjjebekvyWPsfFJmrhwKwmGQyLH","status":404,"time":"2026-03-05T23:17:50+03:00","uri":"/","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.3 Safari/605.1.15"}
{"time":"2026-03-05T23:17:50.8477+03:00","level":"ERROR","prefix":"echo","file":"middleware.go","line":"50","message":"HTTP error: 404 - / - Not Found"}
{"ip":"62.133.61.254","latency":"18.916µs","method":"GET","request_id":"ycHIkvYdncFHiqjiKTwAyVGxujywtsSC","status":404,"time":"2026-03-05T23:17:52+03:00","uri":"/","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.3 Safari/605.1.15"}
{"time":"2026-03-05T23:17:52.509313+03:00","level":"ERROR","prefix":"echo","file":"middleware.go","line":"50","message":"HTTP error: 404 - / - Not Found"}
{"ip":"18.195.23.171","latency":"37.5945ms","method":"DELETE","request_id":"aqCKBHbrlbzONiIWkwyLjWFwMUZsmECR","status":204,"time":"2026-03-05T23:18:02+03:00","uri":"/api/testing/all-data","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"63.296916ms","method":"POST","request_id":"sXglOSjMMiWjKLmEAscrQZXCJjrVxUqY","status":201,"time":"2026-03-05T23:18:02+03:00","uri":"/api/sa/users","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"24.669417ms","method":"DELETE","request_id":"gpRWVMzuCTEGpZmWpefKfRMCpYUbooKu","status":204,"time":"2026-03-05T23:18:02+03:00","uri":"/api/testing/all-data","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"26.404709ms","method":"DELETE","request_id":"pcRDkGEZDoZGRevJJZAcwIzhtVODIIoc","status":204,"time":"2026-03-05T23:18:03+03:00","uri":"/api/testing/all-data","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.797917ms","method":"POST","request_id":"jpLkwcFfuaSrgxbpnQyvQdhjiPQCnLND","status":201,"time":"2026-03-05T23:18:03+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"6.689417ms","method":"PUT","request_id":"aBcYYHwxWDFxLoqbGMYjxhwdtuEQIGwF","status":204,"time":"2026-03-05T23:18:03+03:00","uri":"/api/sa/quiz/questions/fb5c7a3f-08bf-4bd7-9b26-20ae3df10f4a/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.604709ms","method":"POST","request_id":"KDuuxPBTTwIrCFEfVLpjHOeStNMcwHbY","status":201,"time":"2026-03-05T23:18:03+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"2.564667ms","method":"PUT","request_id":"HuKsAiEqqIYHbGmGuWLyAiLFpTmFQKWH","status":204,"time":"2026-03-05T23:18:03+03:00","uri":"/api/sa/quiz/questions/0c317382-7b96-4408-a15a-ab77853b4277/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.55325ms","method":"POST","request_id":"gZmoXPEcwSpSAJDraDwHJjsfkttbLpEf","status":201,"time":"2026-03-05T23:18:03+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"4.669625ms","method":"PUT","request_id":"iHnpXyLFconFZUvrxKzmISrsxrggCpmI","status":204,"time":"2026-03-05T23:18:03+03:00","uri":"/api/sa/quiz/questions/f6578f9a-d2be-452d-b909-7f07490c4c37/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.529042ms","method":"POST","request_id":"TVHiksOjLfcWGzhAdWGFqaPLBjsmgTGe","status":201,"time":"2026-03-05T23:18:04+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"4.435166ms","method":"PUT","request_id":"wXXOhqJyJPhnRKLknqpmbhocIYMwLNZJ","status":204,"time":"2026-03-05T23:18:04+03:00","uri":"/api/sa/quiz/questions/f41cc084-b000-4785-8b33-8bbe49dee476/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"2.826083ms","method":"POST","request_id":"pwvGtAvPQjfkwfyDgIJXuXhzcfHwtXqt","status":201,"time":"2026-03-05T23:18:04+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"5.377334ms","method":"PUT","request_id":"idyXHbHIFaaMISSqAetEPSxXNlXUIQAn","status":204,"time":"2026-03-05T23:18:04+03:00","uri":"/api/sa/quiz/questions/fbdd2a6f-b648-4297-843c-2ab3161fc285/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"63.730292ms","method":"POST","request_id":"tTLtppoVtZkDiqgqGsVKMuFumDxWAjOe","status":201,"time":"2026-03-05T23:18:04+03:00","uri":"/api/sa/users","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"57.015208ms","method":"POST","request_id":"mTpOmpoTeVAfSxGAbrvtTpwwXHSXLiUp","status":201,"time":"2026-03-05T23:18:04+03:00","uri":"/api/sa/users","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"53.732625ms","method":"POST","request_id":"aqIQLjgxHmmFlOrwMFLkGXvvBzkWkpeW","status":200,"time":"2026-03-05T23:18:06+03:00","uri":"/api/auth/login","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"53.933167ms","method":"POST","request_id":"mPmGqwOKdpUPzLaKlBRUsfMwUzDYOrsS","status":200,"time":"2026-03-05T23:18:07+03:00","uri":"/api/auth/login","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.787ms","method":"POST","request_id":"mQnGDvTnDHlBWSVpOfdaAZGNmPhNFSra","status":200,"time":"2026-03-05T23:18:07+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"28.98525ms","method":"POST","request_id":"EAMxDMBzptLtuTdQBfKTGEMazUdPbwVB","status":200,"time":"2026-03-05T23:18:07+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"18.440958ms","method":"POST","request_id":"cPKQWPhjAwwVPZSXLqKDTGFiHCogyiwE","status":200,"time":"2026-03-05T23:18:07+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.515958ms","method":"POST","request_id":"AojEWigJdcFKmAxdzNDpFqiYsgjIDeOx","status":200,"time":"2026-03-05T23:18:07+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.611708ms","method":"POST","request_id":"abyhKJwubedAEjqluPVuzHGdXcuPAMZT","status":200,"time":"2026-03-05T23:18:07+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.100708ms","method":"POST","request_id":"PBPImByqdfYkkKWhTiSbRRNFCBirqohN","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"14.159208ms","method":"POST","request_id":"IRIFnCZyjhATyPnNoVRrgMLlkeeOAtRj","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.037916ms","method":"POST","request_id":"NkeBzbkkESmOdEyMURdmnEZYCasFzxZK","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.904167ms","method":"POST","request_id":"hffPmTStesHmsULGUQRDCCsdOZHsvFJr","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.339375ms","method":"POST","request_id":"WhSiBLPqpJLtrfiffGUVcypzNuRuJLsS","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.745708ms","method":"POST","request_id":"EtVsnqtdhqCQtFlibTlmCBHZOWyWbmbS","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.409041ms","method":"POST","request_id":"ykDtqYGyINWQaYFlUMkTgYIVIxJsHpHn","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.211541ms","method":"POST","request_id":"zbAOrzQeOcILjvKZmtHzYqhfXbATyNLP","status":200,"time":"2026-03-05T23:18:08+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.478208ms","method":"POST","request_id":"VAufGgyjGwSpQNpOoqTRNFuVcquwhXBt","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.836458ms","method":"POST","request_id":"gmSdeYSXvfzxyNlubqaeahkcyKjtrNvy","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.905208ms","method":"POST","request_id":"kMpWQNvztyvjyDEVKEQkoTlkVcSBLCIB","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.323209ms","method":"POST","request_id":"aCTcsZjIAgxqQZtcTKvgtSeWEjlcfjIH","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.522ms","method":"POST","request_id":"QznxODOZjbFnOaExtKyzqfvVtSClQagD","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.702541ms","method":"POST","request_id":"uigyjltJtpPaQgqwqoPhRYsgcvwCYtun","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.2275ms","method":"POST","request_id":"kshcKjtjeldKOanaxdhUMLkLcDJVBgYB","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.686583ms","method":"POST","request_id":"NWjcUSrMpUHvXwQpheSNNheIoRXTLvUe","status":200,"time":"2026-03-05T23:18:09+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.738ms","method":"POST","request_id":"DAveyjAwfXIIYQtJUYeeGJSuAQBndfOo","status":200,"time":"2026-03-05T23:18:10+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"13.095792ms","method":"POST","request_id":"mkOtCLMmcaQGorYpAKuhzRPKpUhJGwFj","status":200,"time":"2026-03-05T23:18:10+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.061708ms","method":"POST","request_id":"rZHwaQARGhOMuqxNfNUlzzdjSullcMmQ","status":200,"time":"2026-03-05T23:18:10+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.666ms","method":"POST","request_id":"UlGFfALFkZwJnlbgdITcAZDdTnBXcPag","status":200,"time":"2026-03-05T23:18:10+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"20.879333ms","method":"POST","request_id":"VGaZxsQdeCZpyZkAGjAZqhOUeSqdUaaY","status":200,"time":"2026-03-05T23:18:10+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.526875ms","method":"POST","request_id":"nVIlceuUEzKipgBggVJqgkmdGMjLcycm","status":200,"time":"2026-03-05T23:18:10+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.382ms","method":"POST","request_id":"WbhmOxGsukvPykAcSTXAqsmVWEvClZJF","status":200,"time":"2026-03-05T23:18:10+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.258792ms","method":"POST","request_id":"yejGkdlykuQhgqndgdmzcSdiBehihlZL","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"14.296167ms","method":"POST","request_id":"uhApjsJhkLZZzRNZcNYdyFJuedpumopo","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.995958ms","method":"POST","request_id":"AjkxWXaxcVEXBLUoWWVzKorGkCIDmyjq","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.056667ms","method":"POST","request_id":"inGSGCAgyEvjfqVanEiGvEAubDXxDcAK","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.929167ms","method":"POST","request_id":"oVpKssFxzkwCtQUMbGjdOceYBHjnJFWe","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.4175ms","method":"POST","request_id":"nZIDitEPapgXgaZTwzqXRXsSoSrcGHVl","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.594125ms","method":"POST","request_id":"eNzvcCuwHlPMAoUIwefwsHjGHMByRBjQ","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.644459ms","method":"POST","request_id":"zpBWDDIjHcGtmpkAmbHsmpHpPNqueEzx","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"2.27475ms","method":"POST","request_id":"qoJGBgcxQohZquQyWTUSUaoXtZTmoWmZ","status":200,"time":"2026-03-05T23:18:11+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"14.15625ms","method":"POST","request_id":"yXQtzvylNfjAAFgCWYqdQQagVMukxhFN","status":200,"time":"2026-03-05T23:18:12+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"20.784875ms","method":"GET","request_id":"wnadyTVNhdwkwutpKMSQcJEIVvmndVhj","status":200,"time":"2026-03-05T23:18:12+03:00","uri":"/api/pair-game-quiz/pairs/my?sortBy=status\u0026sortDirection=asc","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"44.2495ms","method":"DELETE","request_id":"teofBbnyFFQlxgJaDYjrDQKLLmSabbcw","status":204,"time":"2026-03-05T23:18:12+03:00","uri":"/api/testing/all-data","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"5.133542ms","method":"POST","request_id":"lHZvzAcJMyfLgGzvttpmsoUfrkLVCTNV","status":201,"time":"2026-03-05T23:18:12+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.4ms","method":"PUT","request_id":"BRQcTvpFGiUnXaIjEUcoBfgURsfoVxrV","status":204,"time":"2026-03-05T23:18:12+03:00","uri":"/api/sa/quiz/questions/fb40f0c6-c425-4d72-a42c-6f4120e45637/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.33425ms","method":"POST","request_id":"IhoPRPCBOOEikbFfsSoYYJQiReARPaFv","status":201,"time":"2026-03-05T23:18:12+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.740334ms","method":"PUT","request_id":"IfDGMqrbmniyvbSWXUpIhSDpyXcsJLmQ","status":204,"time":"2026-03-05T23:18:12+03:00","uri":"/api/sa/quiz/questions/fb092ada-2fc2-42bc-922e-6c919a63e88c/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.122375ms","method":"POST","request_id":"fVKYhLUXJnwOiFOxirZEZWOAypwKlijf","status":201,"time":"2026-03-05T23:18:13+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.970041ms","method":"PUT","request_id":"yRAIJoBVbzrQpCSjnJxVfueMTEviIuOP","status":204,"time":"2026-03-05T23:18:13+03:00","uri":"/api/sa/quiz/questions/71ac0f15-b1a9-407b-bb5f-f2d2b8f2fd60/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"2.279417ms","method":"POST","request_id":"fNDTQcqQMOvvHlUrquvUrqJsxmQIKVcC","status":201,"time":"2026-03-05T23:18:13+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.988416ms","method":"PUT","request_id":"GzxYVNHVqAQqEAzzyxwskIGvAXHSZIVA","status":204,"time":"2026-03-05T23:18:13+03:00","uri":"/api/sa/quiz/questions/21814aa6-6595-4d6b-ba94-545c97acbb87/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"2.81525ms","method":"POST","request_id":"uWwmVFaXzbZsDuKifIxeivrHQNoNWlxJ","status":201,"time":"2026-03-05T23:18:13+03:00","uri":"/api/sa/quiz/questions","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"3.737833ms","method":"PUT","request_id":"GudHsqOWmOquftqwNPoHZkCoMKFkArrQ","status":204,"time":"2026-03-05T23:18:13+03:00","uri":"/api/sa/quiz/questions/b29dc337-8411-4fd8-8f05-797979027980/publish","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"61.424292ms","method":"POST","request_id":"qGCtCttMxtmUDyfTZGnOwvoPvjpKgluR","status":201,"time":"2026-03-05T23:18:13+03:00","uri":"/api/sa/users","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"57.378083ms","method":"POST","request_id":"aKOgwwdxgcMIHDuNzPGgrAtLVqxDCLSk","status":201,"time":"2026-03-05T23:18:14+03:00","uri":"/api/sa/users","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"82.057416ms","method":"POST","request_id":"buONTPPeEDkhrRwMwYzDwjAhzuIGFDLx","status":201,"time":"2026-03-05T23:18:14+03:00","uri":"/api/sa/users","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"59.260542ms","method":"POST","request_id":"uoFCzSLvMSCatQCVwystPihOzcgvSHng","status":201,"time":"2026-03-05T23:18:14+03:00","uri":"/api/sa/users","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"55.722084ms","method":"POST","request_id":"jwdbOKkALjywOYKhPpvAyENNmpDLFFre","status":200,"time":"2026-03-05T23:18:15+03:00","uri":"/api/auth/login","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"54.528916ms","method":"POST","request_id":"tyCDmzExmCqwaEeMLDhYYlCtwsJeWrDw","status":200,"time":"2026-03-05T23:18:16+03:00","uri":"/api/auth/login","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"56.025959ms","method":"POST","request_id":"bCTmCRXtvDcMjZghEpBFJJAyVJTpTrFW","status":200,"time":"2026-03-05T23:18:18+03:00","uri":"/api/auth/login","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"56.88425ms","method":"POST","request_id":"nhrNWJwAiipGSAVNYtJkNQbxFuHVdpiv","status":200,"time":"2026-03-05T23:18:19+03:00","uri":"/api/auth/login","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"13.112666ms","method":"POST","request_id":"GaFNKplxiLHTVBZcMpzuoudBjOLRpPRP","status":200,"time":"2026-03-05T23:18:19+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"17.4765ms","method":"POST","request_id":"WBhYaLCGaYjpxCYfZYAQgnBnhrvMQKGL","status":200,"time":"2026-03-05T23:18:19+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.009333ms","method":"POST","request_id":"NKxIlKzDwCnEZgEBHAorevMowUOsmqgB","status":200,"time":"2026-03-05T23:18:19+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.508541ms","method":"POST","request_id":"xmrnSXAzxGawZiPlHPFCCoBfnIrLuyoa","status":200,"time":"2026-03-05T23:18:19+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.294541ms","method":"POST","request_id":"BVgqddBCwYRorgNUitcgoaFYfrPiHpMe","status":200,"time":"2026-03-05T23:18:19+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.648375ms","method":"POST","request_id":"hmelRssUbUNolApcAtTTzEfKMTIhRBgL","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.299791ms","method":"POST","request_id":"NaGWBmGscbJLarChCAPZQntloBJLJaiZ","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.398709ms","method":"POST","request_id":"ibOLYshvfphzPOwYjqGiOPHDaXiTOhVt","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.138ms","method":"POST","request_id":"JcSrfNalLNREbyhMXodtGtJDlHscCdZr","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.778291ms","method":"POST","request_id":"nznFwgMYObiyOQBQGzTwMwkkFIcXwEfK","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.235666ms","method":"POST","request_id":"hhPXyMjZrWdyDVbReMEqwBZhOcUbNRTN","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"13.92875ms","method":"POST","request_id":"aSGjbyFTiOVTiXPJnJQJjFflQPBUKzoV","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.7185ms","method":"POST","request_id":"zZBNgJKLdGfrSnBnDgPQmCYHagvEsdIB","status":200,"time":"2026-03-05T23:18:20+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"17.842958ms","method":"POST","request_id":"AmvwRNdlxOOvQLYFhsBZQpaiKFvYBgxY","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.751042ms","method":"POST","request_id":"uqMTysAcQDosNoPkvcTvIwTeFXVsBHVl","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"13.18475ms","method":"POST","request_id":"xfisgSAXlpJEFHVgoHFFFeswFIHsNxIa","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.943667ms","method":"POST","request_id":"HuUAMgTxefxbsmuFdJcpkOFVnuICLCiR","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.855792ms","method":"POST","request_id":"qKfLTKZioWsAAguJJzZzIPqeAlApLAJE","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.048833ms","method":"POST","request_id":"gATqTzISBuZbSYFdJxKzFiCsubOzioMC","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.604166ms","method":"POST","request_id":"stZsmYMHGFUnmXgAlqpVqtjwwPIyDVnU","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.259209ms","method":"POST","request_id":"qbEWeXPOygHnkinaNqdHsCQtHxszUIfI","status":200,"time":"2026-03-05T23:18:21+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.5905ms","method":"POST","request_id":"kDUkjNpLEQFmLxCpCelZKtTmzLJDcxmS","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.2455ms","method":"POST","request_id":"PPlxDffgoVmbKEcFmkEykZQasfgAspLk","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"4.966167ms","method":"POST","request_id":"dVCZCaKkSgOctZrkVaQJmrpZuaFpGUTl","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.803375ms","method":"POST","request_id":"UezpWskyGsPHMsbBMwGmYbpIqOrNmHJU","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"14.885666ms","method":"POST","request_id":"TdfHqtkDKIBxQIECtsfJACnItQIRlCTN","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.356792ms","method":"POST","request_id":"DdJcMsBzjdcAYUlLmRxDkPUDDqvWjMHQ","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.899417ms","method":"POST","request_id":"QFUKYwBGxuFhOMWxXJngswwvzolNyBkd","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.070708ms","method":"POST","request_id":"LNhLaGeoyeASWtNQkijFqLmqTsTCbUzj","status":200,"time":"2026-03-05T23:18:22+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.035084ms","method":"POST","request_id":"bahmRriBjPmbOXZjLvMhPdKftpDiIryI","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.56875ms","method":"POST","request_id":"cKQVSNnJedqJPatxUnjHAHwbDisByFrq","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.001ms","method":"POST","request_id":"KkKnBCbbrEwykPKIrRWUyCxpqaHTtWkl","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.974042ms","method":"POST","request_id":"LyKFHRfeDQInHjiEPPzyHaiuEvsItKBP","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.117917ms","method":"POST","request_id":"KFEZRnawyFWIPgcdpgjPgsrNtDMKJFqe","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.9735ms","method":"POST","request_id":"JqIJCaocUUsWfoXMPGGYDEENOcofyLzm","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.347916ms","method":"POST","request_id":"ceFPjpxVZOUQnLhrfIYMlipUIQQTcFut","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.938959ms","method":"POST","request_id":"HOKoXsfdIVLeGPxbULmKLenuWIDdYeXf","status":200,"time":"2026-03-05T23:18:23+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"16.53075ms","method":"POST","request_id":"zwnzfvhNnoDLXWGuOzPxyZizMrrXuMxg","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.454917ms","method":"POST","request_id":"GjdHYNIgDwYQbBKzTFxXuFvppSuEugEj","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"14.536625ms","method":"POST","request_id":"RviwCJWWUqYkUlslYrdpgUDxPlyjHqXn","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.981917ms","method":"POST","request_id":"gcXvZGkGtuAfxhyafNrXHaxHizSgilPr","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.642709ms","method":"POST","request_id":"KcYDsjcYxQPeENkJWhhqCpdvkfMZlYBZ","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.382291ms","method":"POST","request_id":"JSNJvRghXqZIPzFYmXQRdwQZROOlpXxU","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.63025ms","method":"POST","request_id":"ysUfGKfqhbCABoiqeosfuQkEpPkZdIZy","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.109583ms","method":"POST","request_id":"KVAOCQHZgYDpXIRDyOWqvpMwsnpccpEx","status":200,"time":"2026-03-05T23:18:24+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.375042ms","method":"POST","request_id":"jxpdHIMUqyeEYTwZmFPMmfyaiUmIhVbW","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.013375ms","method":"POST","request_id":"uimJqhsfALnKeUxLiGkamhgmNIGFWuqc","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.487792ms","method":"POST","request_id":"JbposqLAaDATJFMSLkQZGspgkfLiteWH","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"5.3765ms","method":"POST","request_id":"MgCWvevLAhPhkvLLXtVaexeaHDqhuzdF","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.098375ms","method":"POST","request_id":"nkyBRfzgzxELRWzYyVDUACpPIJElaZeo","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.334666ms","method":"POST","request_id":"lfHsekLTyjzwJDlScqCLWCwHJOlNCGIf","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.249625ms","method":"POST","request_id":"aRYWZWRhXTcMSeJoArndcygsqmEsoJIr","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.501916ms","method":"POST","request_id":"SUNGqBAklJLATAuDUDtAPOTPQWReMQnC","status":200,"time":"2026-03-05T23:18:25+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.769417ms","method":"POST","request_id":"UjxgjRIVECUcJKqjEzoNgctExeohmnXA","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.424ms","method":"POST","request_id":"ZFZBFnuSCxocgzrxTrZqbOnqZwDQfsdZ","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.879709ms","method":"POST","request_id":"GyfngkkpXOttApxNdrrjOWSYvsNnDngc","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.212083ms","method":"POST","request_id":"JFnEouCMZzKLGXQmUjgVYCUksTYpMldX","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.5515ms","method":"POST","request_id":"VMpYithWIDYddHvELaNlzPjNbxaDdLFi","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.042959ms","method":"POST","request_id":"VqnAskQOdyPzrGQXNVYCLMcWfOCJmkYx","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.246791ms","method":"POST","request_id":"CSFzrMAGcCEAZOzTJaJAuxFHtAbOVmmN","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"5.736042ms","method":"POST","request_id":"fcvKpfqxvhxcYjHlbihdxihPnxgGgOnf","status":200,"time":"2026-03-05T23:18:26+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"13.082791ms","method":"POST","request_id":"futyanbcGFWknfCSoDOOrAShYHslIXeL","status":200,"time":"2026-03-05T23:18:27+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.995958ms","method":"POST","request_id":"KHagTfzaGewWRkWNZMYnmZhulWHxBMQa","status":200,"time":"2026-03-05T23:18:27+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"14.201292ms","method":"POST","request_id":"pWiwZNEZfKpMcYDkeyWlJaatihLoNPpy","status":200,"time":"2026-03-05T23:18:27+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.5025ms","method":"POST","request_id":"MnStOPGnAbXvCgZxpcuihbufZsvuudIq","status":200,"time":"2026-03-05T23:18:27+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.365292ms","method":"POST","request_id":"keznvKlwEezIzwPvmwPUdjdzlrNxHAgD","status":200,"time":"2026-03-05T23:18:27+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.035083ms","method":"POST","request_id":"MFaaLvRJZgNhmxkZDvIxcdwyJywsjTjd","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.69525ms","method":"POST","request_id":"tuOwkhBRJSPSsRcjVpMKTTxvoOOWmzvo","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.219667ms","method":"POST","request_id":"lfkknPHexRXQLiSeQoXfbZaWrXLgsnof","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.616334ms","method":"POST","request_id":"eMunQaZawTAuiTHodCrPZNONPgHYIqNH","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.748833ms","method":"POST","request_id":"TtMHWlgPfHqaougYTQuLViqszZReCmKZ","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.779125ms","method":"POST","request_id":"UdrGWfrILwmNuCRqUAQNpVOjvktlKorP","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"6.720708ms","method":"POST","request_id":"BQZxidzZmMgcNTZbnbYrtLQwMzUcPBBW","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"13.725958ms","method":"POST","request_id":"wOClpcvbSzVXFKIpyXPhpumNoyXZODuu","status":200,"time":"2026-03-05T23:18:28+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.901583ms","method":"POST","request_id":"HjWlyyveJTMVcNdvEemcBLZxmbbJcFYj","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.006042ms","method":"POST","request_id":"MTYuEELzySKgcThPjgxcegZXMPfyeclp","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.354917ms","method":"POST","request_id":"NqqZmuLDArhHmAOrVnynawfZTKOcZJkN","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.866042ms","method":"POST","request_id":"QWycXdUaSflGZISpTLVYnssVnEqSraEN","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.933458ms","method":"POST","request_id":"yeZQYvmFtBoMBYXLTdlyQzzdGFoQuAjH","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"12.56225ms","method":"POST","request_id":"PKIEbuRffCLNVuboiQCjgueHwpjpdvtg","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.086041ms","method":"POST","request_id":"KslRaumwChvCNTaWQCgMUwKPUyZbmonV","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.478334ms","method":"POST","request_id":"hjiTShclQWNVCOVCVztlwXHFdFPnJoZR","status":200,"time":"2026-03-05T23:18:29+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.355958ms","method":"POST","request_id":"ejhbhEwcgjQTkGkXevHQDuBZEcWPXrLU","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.839833ms","method":"POST","request_id":"YfmMRqwnVcjPELYsRQipoFUtJRNlZbGA","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"5.804292ms","method":"POST","request_id":"aYWaKwxcJMfvYcfEnGFbLlLQfTtXPpyh","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"16.225167ms","method":"POST","request_id":"nAvKkFPdAYgAValeHECRhBtrtbdzeoVo","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/connection","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.882458ms","method":"POST","request_id":"NNHNkKubJtPgKtengfroMLifKGoCiwgG","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.8095ms","method":"POST","request_id":"WjHmKWdtdCsquskOBKrOHFpCtjJZppag","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.879209ms","method":"POST","request_id":"mHHzqGmtxcEwFOhgmRoluTcQGwRdxEvE","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.046833ms","method":"POST","request_id":"SgXmmJUuJqWoqTzQDfBGIkiKDisgDEUp","status":200,"time":"2026-03-05T23:18:30+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"7.954666ms","method":"POST","request_id":"dMfsDnHpluoihhmSjtZtlFuXmzYmdFuG","status":200,"time":"2026-03-05T23:18:31+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"8.626209ms","method":"POST","request_id":"NEKmpHLJdCNwoQKCitbuZnmZGftvSOLA","status":200,"time":"2026-03-05T23:18:31+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.24925ms","method":"POST","request_id":"wMjfFZUpztframEzeNIAIliutenpyFNh","status":200,"time":"2026-03-05T23:18:31+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"10.6695ms","method":"POST","request_id":"mXXGXIudfJkhfkHAAWHunhLILlgNfWix","status":200,"time":"2026-03-05T23:18:31+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.086292ms","method":"POST","request_id":"UuKxnpcltmRvwJSqgtSnhhcOVJtTUBDM","status":200,"time":"2026-03-05T23:18:31+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"11.280458ms","method":"POST","request_id":"qagiGtfZawZOROymQPRdpvhTudIrerhT","status":200,"time":"2026-03-05T23:18:31+03:00","uri":"/api/pair-game-quiz/pairs/my-current/answers","user_agent":"axios/1.12.0"}
{"ip":"18.195.23.171","latency":"9.221833ms","method":"GET","request_id":"DdfwCKumSUFQstTIPgMCCSKOWZRXtLJc","status":200,"time":"2026-03-05T23:18:31+03:00","uri":"/api/pair-game-quiz/users/my-statistic","user_agent":"axios/1.12.0"}
