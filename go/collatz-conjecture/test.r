#!/usr/bin/env Rscript

# Collatz Conjecture (Гипотеза Коллатца)
# Скрипт:
#  - реализует функции для получения цепочки Коллатца и количества шагов до 1 (stopping time)
#  - выполняет расчет для n = 1..N_MAX (вверху скрипта задаётся `N_MAX` или можно передать первым аргументом командной строки)
# Параметры (можно переопределить первым аргументом командной строки)
N_MAX <- 1000L
args <- commandArgs(trailingOnly = TRUE)
if (length(args) >= 1) {
  maybeN <- suppressWarnings(as.integer(args[1]))
  if (!is.na(maybeN) && maybeN > 0) {
    N_MAX <- maybeN
  }
}
cat(sprintf("Используется N_MAX = %d\n", N_MAX))
# Базы для меток по Y (напр., c(1,2,5)); можно поменять в начале скрипта
YTICK_BASES <- c(1, 2, 5)
if (!requireNamespace("ggplot2", quietly = TRUE)) {
  message("Пакет 'ggplot2' не найден — пытаюсь установить...")
  install.packages("ggplot2", repos = "https://cloud.r-project.org")
  if (!requireNamespace("ggplot2", quietly = TRUE)) {
    stop("Не удалось установить пакет 'ggplot2'. Установите его вручную: install.packages('ggplot2')")
  }
}
suppressPackageStartupMessages(library(ggplot2))
#  - строит график зависимости количества шагов от начального n и сохраняет его в PNG

# Функция: возвращает всю цепочку Коллатца, начиная с n и заканчивая 1
collatz_sequence <- function(n, max_iter = 10000L) {
  if (!is.numeric(n) || length(n) != 1 || n <= 0 || floor(n) != n) {
    stop("n must be a positive integer")
  }
  seq <- as.numeric(n)
  iter <- 0L
  while (tail(seq, 1) != 1 && iter < max_iter) {
    last <- tail(seq, 1)
    if (last %% 2 == 0) {
      seq <- c(seq, last / 2)
    } else {
      seq <- c(seq, 3 * last + 1)
    }
    iter <- iter + 1L
  }
  if (iter >= max_iter && tail(seq, 1) != 1) {
    warning("Reached max_iter without reaching 1")
  }
  seq
}

# Функция: возвращает число шагов до 1 (stopping time) с мемоизацией (быстро для последовательных вызовов)
collatz_steps <- (function() {
  cache <- new.env(parent = emptyenv())
  cache[["1"]] <- 0L
  function(n, max_iter = 1e6) {
    if (!is.numeric(n) || length(n) != 1 || n <= 0 || floor(n) != n) {
      stop("n must be a positive integer")
    }
    key <- as.character(as.integer(n))
    if (exists(key, envir = cache, inherits = FALSE)) {
      return(as.integer(cache[[key]]))
    }

    m <- n
    path <- integer(0)
    iter <- 0L
    while (!exists(as.character(as.integer(m)), envir = cache, inherits = FALSE) &&
           m != 1 && iter < max_iter) {
      path <- c(path, as.integer(m))
      if (m %% 2 == 0) {
        m <- m / 2
      } else {
        m <- 3 * m + 1
      }
      iter <- iter + 1L
    }

    if (iter >= max_iter &&
        !exists(as.character(as.integer(m)), envir = cache, inherits = FALSE) &&
        m != 1) {
      warning(sprintf("Reached max_iter=%d without converging for initial n=%d", max_iter, n))
    }

    base <- 0L
    if (exists(as.character(as.integer(m)), envir = cache, inherits = FALSE)) {
      base <- as.integer(cache[[as.character(as.integer(m))]])
    }

    if (length(path) > 0) {
      for (i in seq_len(length(path))) {
        val <- path[length(path) - i + 1]
        base <- base + 1L
        cache[[as.character(as.integer(val))]] <- base
      }
    }

    as.integer(cache[[key]])
  }
})()

# --- Выполнение расчета для n = 1..N_MAX ---
N_MAX <- as.integer(N_MAX)
if (is.na(N_MAX) || N_MAX < 1L) stop("N_MAX must be a positive integer >= 1")
ns <- 1:N_MAX
steps <- vapply(ns, collatz_steps, integer(1L))

# Статистика
max_idx <- which.max(steps)
max_n <- ns[max_idx]
max_steps <- steps[max_idx]
cat(sprintf("Максимальное число шагов: %d (n = %d)\n", max_steps, max_n))

top10_idx <- order(steps, decreasing = TRUE)[1:10]
top10 <- data.frame(n = ns[top10_idx], steps = steps[top10_idx])
cat("Топ-10 начальных n по количеству шагов:\n")
print(top10)

# --- Построение графика зависимости steps(n) через ggplot2 и сохранение (профессионально) ---
df_steps <- data.frame(n = ns, steps = steps)
out_steps_png <- sprintf("collatz_steps_1_%d.png", N_MAX)

# pick top 3 for annotation
top3 <- head(top10, 3)

p_steps <- ggplot(df_steps, aes(x = n, y = steps)) +
  geom_point(alpha = 0.6, size = 1.1, color = "#2c7fb8") +
  geom_smooth(method = "loess", se = TRUE, color = "#d95f0e", fill = "#fdd49e", linewidth = 0.9, span = 0.25) +
  geom_point(data = subset(df_steps, n == max_n), aes(x = n, y = steps), color = "#d73027", size = 3) +
  geom_text(data = top3, aes(label = sprintf("n=%d\n%d", n, steps)), vjust = -0.6, size = 3.2, color = "#444444") +
  labs(title = sprintf("Collatz: stopping time (n = 1..%d)", N_MAX),
       subtitle = sprintf("max: n=%d, steps=%d", max_n, max_steps),
       x = "Исходное число (n)",
       y = "Число шагов до 1") +
  theme_minimal(base_size = 14) +
  theme(plot.title = element_text(face = "bold", size = 16, hjust = 0.5),
        plot.subtitle = element_text(size = 12, hjust = 0.5, color = "#555555"),
        axis.title = element_text(size = 12),
        axis.text = element_text(size = 10),
        panel.grid.major = element_line(color = "#e9e9e9"),
        panel.grid.minor = element_blank())

# Save at higher quality
ggsave(out_steps_png, plot = p_steps, width = 12, height = 6, dpi = 300)
cat(sprintf("GGplot (профессиональный) сохранён в файле %s\n", out_steps_png))

# (Scatter-plot удалён по запросу)

# Если запуск интерактивный, показываем график сразу
if (interactive()) {
  print(p_steps)
  print(p)
}

# Дополнительно: выводим цепочку для n с максимальным числом шагов и сохраняем её график
seq_max_n <- collatz_sequence(max_n)
cat(sprintf("Для n=%d: длина цепочки = %d (шагов), максимальное значение в цепочке = %d\n",
            max_n, length(seq_max_n) - 1, max(seq_max_n)))

# Построение через ggplot2 с лог-шкалой по Y и метками на 1-2-5 * 10^k
seq_df <- data.frame(step = seq_along(seq_max_n) - 1L, value = seq_max_n)

min_val <- min(seq_max_n)
max_val <- max(seq_max_n)
min_exp <- floor(log10(min_val))
max_exp <- ceiling(log10(max_val))
# Используем YTICK_BASES для генерации кандидатов (базы * 10^e)
cand_ticks <- unlist(lapply(min_exp:max_exp, function(e) (10^e) * YTICK_BASES))
yticks <- sort(unique(cand_ticks[cand_ticks >= min_val & cand_ticks <= max_val]))
if (length(yticks) == 0) {
  # fallback к степеням 10 если ничего не найдено
  yticks <- 10^(min_exp:max_exp)
  yticks <- yticks[yticks >= min_val & yticks <= max_val]
}

# Профессиональный лог-график цепочки (ggplot)
max_seq_step <- which.max(seq_max_n) - 1L
max_seq_val <- max(seq_max_n)

p <- ggplot(seq_df, aes(x = step, y = value)) +
  geom_line(color = "#1f77b4", linewidth = 0.9) +
  geom_point(size = 1.2, color = "#1f77b4", alpha = 0.9) +
  geom_point(data = data.frame(step = max_seq_step, value = max_seq_val), aes(x = step, y = value), color = "#d73027", size = 3) +
  geom_text(data = data.frame(step = max_seq_step, value = max_seq_val), aes(x = step, y = value, label = sprintf("max=%d", value)), vjust = -1, size = 3.2, color = "#333333") +
  scale_y_log10(breaks = yticks, labels = scales::label_number(accuracy = 1, big.mark = " ")) +
  labs(x = "шаг", y = "значение (лог шкала)", title = sprintf("Цепочка Коллатца для n = %d (лог шкала)", max_n),
       subtitle = sprintf("максимум в цепочке: %d", max_seq_val)) +
  theme_minimal(base_size = 14) +
  theme(plot.title = element_text(face = "bold", size = 16, hjust = 0.5),
        plot.subtitle = element_text(size = 12, hjust = 0.5, color = "#555555"),
        axis.title = element_text(size = 12),
        axis.text = element_text(size = 10),
        panel.grid.major = element_line(color = "#e9e9e9"),
        panel.grid.minor = element_blank())

# Save with higher resolution
ggsave(filename = sprintf("collatz_sequence_n_%d.png", max_n), plot = p, width = 10, height = 5, dpi = 300)
cat(sprintf("GGplot (профессиональный) сохранён в файле collatz_sequence_n_%d.png\n", max_n))
