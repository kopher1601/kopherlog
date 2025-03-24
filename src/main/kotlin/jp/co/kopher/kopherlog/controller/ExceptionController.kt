package jp.co.kopher.kopherlog.controller

import jp.co.kopher.kopherlog.response.ErrorResponse
import org.slf4j.LoggerFactory
import org.springframework.http.HttpStatus
import org.springframework.web.bind.MethodArgumentNotValidException
import org.springframework.web.bind.annotation.ExceptionHandler
import org.springframework.web.bind.annotation.ResponseStatus
import org.springframework.web.bind.annotation.RestControllerAdvice

@RestControllerAdvice
class ExceptionController {

    val log = LoggerFactory.getLogger(ExceptionController::class.java)

    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ExceptionHandler
    fun invalidRequestHandler(e: MethodArgumentNotValidException): ErrorResponse? {
        log.info("invalidRequestHandler : ${e.message}")
        val response = ErrorResponse("400", "リクエストに誤りがあります。")
        e.fieldErrors.forEach { fieldError ->
            response.addValidation(fieldError.field, fieldError.defaultMessage!!)
        }

        return response
    }
}