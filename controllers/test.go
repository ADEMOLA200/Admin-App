package controllers

// func (a *HackathonController) DeleteProfile(c *gin.Context) {
// 	uid := c.Param("uid")

// 	err := a.hackathonSvc.DeleteProfileByUID(c, uid)
// 	if err!= nil {
// 		c.JSON(err.StatusCode, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Profile deleted successfully",
// 		"success": true,
// 	})
// }

// func (h *HackathonService) DeleteProfileByUID(c *gin.Context, uid string) *common.RestErr {
// 	var profile models.HackathonProfile
// 	err := h.hackathonRepo.GetProfileByUID(uid, &profile)
// 	if err!= nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return h.restErr.NotFound("Profile not found")
// 		}
// 		return h.restErr.ServerError("Failed to get profile")
// 	}

// 	err = h.hackathonRepo.DeleteProfileByID(profile.ID)
// 	if err!= nil {
// 		return h.restErr.ServerError("Failed to delete profile")
// 	}

// 	return nil
// }

// func (h *HackathonRepository) GetProfileByUID(uid string, profile *models.HackathonProfile) error {
// 	result := h.db.Where("uid =?", uid).First(profile)
// 	return result.Error
// }

// func (h *HackathonRepository) DeleteProfileByID(id int) error {
// 	result := h.db.Delete(&models.HackathonProfile{}, id)
// 	if result.Error!= nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		return gorm.ErrRecordNotFound
// 	}
// 	return nil
// }