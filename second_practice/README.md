# ИНСТРУКЦИЯ К ЗАПУСКУ И ИСПОЛЬЗОВАНИЮ:
1. Установить golang (используемая верси 1.18.1), добавить в переменные окружения
2. Склонировать репозиторий
3. Перейти в директорию second_practice (основная прога)
4. запустить приложение: команда go run ./app/main.go

# БАГИ
1. Неправильная handle-функция для эндпоинта - GET /api/historicalCurrenciesByPeriod
2. Неправильный ключ apikey в хедере для обращения к Currency Layer API в методе --GetSupportedCurrencies-- сервиса requester