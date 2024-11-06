package jp.co.kopher1601.topu.recruitment.service

import jp.co.kopher1601.topu.recruitment.service.dto.PostRecruitment

interface RecruitmentService {
    fun post(postRecruitment: PostRecruitment)
}