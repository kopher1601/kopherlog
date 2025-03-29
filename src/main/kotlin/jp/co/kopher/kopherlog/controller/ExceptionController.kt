package jp.co.kopher.kopherlog.controller

import jp.co.kopher.kopherlog.exception.KopherlogException
import jp.co.kopher.kopherlog.response.ErrorResponse
import org.slf4j.LoggerFactory
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
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
        val response = ErrorResponse(
            _code = "400",
            _message = "リクエストに誤りがあります。"
        )
        e.fieldErrors.forEach { fieldError ->
            response.addValidation(fieldError.field, fieldError.defaultMessage!!)
        }

        return response
    }

    @ExceptionHandler(KopherlogException::class)
    fun postNotFound(ex: KopherlogException): ResponseEntity<ErrorResponse> {
        log.info("postNotFound : ${ex.message}")
        val response = ErrorResponse(
            _code = ex.code.toString(),
            _message = ex.message!!,
            _validation = ex.validation
        )

        return ResponseEntity.status(ex.code).body(response)
    }
}