package jp.co.kopher1601.topu.recruitment.service

import jp.co.kopher1601.topu.recruitment.service.dto.PostRecruitment
import jp.co.kopher1601.topu.recruitment.service.dto.RecruitmentResponse

interface RecruitmentService {
    fun post(postRecruitment: PostRecruitment)
    fun getRecruitment(recruitmentId: Long): RecruitmentResponse
}