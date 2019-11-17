local letSunDisappearDone = false
local letNewStarAppearDone = false

function letSunDisappear()
  if getSteps() >= 365 * 1 and not letSunDisappearDone then
    if getBodyById("sun") ~= nil then
      removeBodyById("sun")
      letSunDisappearDone = true
    end
  end

  if getSteps() >= 365 * 2 and not letNewStarAppearDone then
    local position = createPoint3D(0, 0, 0)
    local velocity = createVector3D(0, 0, 0)
    local body = createBody("sun2", "new sun", 1.98892e30, 1, position, velocity)
    addBodyToScene(body, 255, 255, 255, 255)
    letNewStarAppearDone = true
  end
end
