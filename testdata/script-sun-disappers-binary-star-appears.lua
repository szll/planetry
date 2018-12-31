local swapDone = false

function letSunDisappear()
  local steps = getSteps()

  if steps == 364 or steps == 366 then
    setPaused(true)
  end

  if steps >= 365 and not swapDone then
    local sun = getBodyById("sun")
    
    -- Remove Sun
    if getBodyById("sun") ~= nil then
      removeBodyById("sun")
      letSunDisappearDone = true
    end

    local mass = sun.Mass / 2
    local velocity = 33500
    local distance = 0.1 * AU

    -- Add a binary star where the sun was, both stars should be stable ...
    local alpha = createBody("alpha", "Alpha", mass, 1, createPoint3D(0, distance, 0), createVector3D(velocity, 0, 0))
    addBodyToScene(alpha, 255, 0, 255, 255)
    print("created alpha")
    
    local beta = createBody("beta", "Beta", mass, 1, createPoint3D(0, -distance, 0), createVector3D(-velocity, 0, 0))
    addBodyToScene(beta, 255, 0, 255, 255)
    print("created beta")

    swapDone = true
  end
end
