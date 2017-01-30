import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from "../api/api.service";
import { ClusterModel } from "../api/model/ClusterModel";
//import { ClusterEntity } from "../api/entitiy/ClusterEntity";
import { NodeEntity } from "../api/entitiy/NodeEntity";

@Component({
  selector: 'kubermatic-cluster',
  templateUrl: './cluster.component.html',
  styleUrls: ['./cluster.component.scss']
})
export class ClusterComponent implements OnInit {

  constructor(private route: ActivatedRoute, private api: ApiService) {}

  public routingParams: any;
  public nodes: any;
  public cluster: any;

  ngOnInit() {

    this.route.params.subscribe(params => {
      this.routingParams = new ClusterModel(params['seedDcName'], params['clusterName']);
      this.getCluster(this.routingParams);
      this.getNodes(this.routingParams);
    })
  }

  getCluster(clusterModel){
    this.api.getCluster(clusterModel).subscribe(result => {
      this.cluster = result;
    });
  }

  getNodes(clusterModel){
    this.api.getClusterNodes(clusterModel).subscribe(result => {
        this.nodes = result;
    })
  }

  public addNodes(clusterModel): void {
    //this.api.createClusterNode(clusterModel, /*  node  */);

  }
}

