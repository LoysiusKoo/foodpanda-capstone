@POST("dishes")
suspend fun getFiltered(
	@Url url: String, 
	@Query("uri1") uri1: String, 
	@Query("uri2") uri2: String, 
	@Query("uri3") uri3: Float, 
	@Query("uri4") uri4: Float, 
	@Query("uri5") uri5: Float, 
	@Body requestBody: RequestBody
	): Response<YourResponseModel>