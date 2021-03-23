require "mustache"
class Simple < Mustache

  $inputs={"httpApplicationRoutingDomain"=>false , "reuseACR"=> true}

    def registryUrl
      return "ww.w.w.w"
    end

  def reuseACR
    true
  end

  def httpApplicationRoutingDomain
    return false
  end
  def projects
    return ["p1","p2","p3","p4"]
  end
end

Simple.template_file ="template.yml"

file = File.open("template-output.yml","w")
file.write(Simple.render)

