package jp.co.kopher1601.topu.recruitment.controller

import jp.co.kopher1601.topu.recruitment.controller.dto.PostRecruitmentRequest
import jp.co.kopher1601.topu.recruitment.service.dto.PostRecruitment
import jp.co.kopher1601.topu.recruitment.service.RecruitmentService
import jp.co.kopher1601.topu.recruitment.service.dto.RecruitmentResponse
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*

@RestController
class RecruitmentController @Autowired constructor(
    private val recruitmentService: RecruitmentService,
) {

    private val log = LoggerFactory.getLogger(this::class.java)

    @ResponseStatus(HttpStatus.CREATED)
    @PostMapping("/recruitments")
    fun post(@RequestBody request: PostRecruitmentRequest) {
        recruitmentService.post(PostRecruitment.from(request))
    }

    @GetMapping("/recruitments/{recruitmentId}")
    fun getRecruitment(@PathVariable recruitmentId: Long): RecruitmentResponse {
        return recruitmentService.getRecruitment(recruitmentId)
    }
}