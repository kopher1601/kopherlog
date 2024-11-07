package jp.co.kopher1601.topu

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.data.jpa.repository.config.EnableJpaAuditing

@EnableJpaAuditing
@SpringBootApplication
class TopuApplication

fun main(args: Array<String>) {
    runApplication<TopuApplication>(*args)
}
