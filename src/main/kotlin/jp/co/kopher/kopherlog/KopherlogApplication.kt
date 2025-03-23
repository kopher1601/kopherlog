package jp.co.kopher.kopherlog

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class KopherlogApplication

fun main(args: Array<String>) {
    runApplication<KopherlogApplication>(*args)
}
