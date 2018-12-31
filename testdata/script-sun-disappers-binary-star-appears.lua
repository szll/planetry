local letSunDisappearDone = false
local letNewStarAppearDone = false

function letSunDisappear()
  if getSteps() >= 365 and not letSunDisappearDone then
    if getBodyById("sun") ~= nil then
      removeBodyById("sun")
      letSunDisappearDone = true
    end
  end

  if getSteps() >= 366 and not letNewStarAppearDone then
    local mass = 1.98892e30 / 2
    local velocity = 33500
    local distance = 0.1 * AU

    local alpha = createBody("alpha", "Alpha", mass, 1, createPoint3D(0, distance, 0), createVector3D(velocity, 0, 0))
    addBodyToScene(alpha, 255, 255, 0, 255)
    print("created alpha")
    
    local beta = createBody("beta", "Beta", mass, 1, createPoint3D(0, -distance, 0), createVector3D(-velocity, 0, 0))
    addBodyToScene(beta, 255, 0, 255, 255)
    print("created beta")

    letNewStarAppearDone = true
  end

  if getSteps() == 368 or getSteps() == 400 then
    setPaused(true)
  end
end
