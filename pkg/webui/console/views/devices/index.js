// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { Switch, Route } from 'react-router-dom'
import { connect } from 'react-redux'

import Breadcrumb from '../../../components/breadcrumbs/breadcrumb'
import { withBreadcrumb } from '../../../components/breadcrumbs/context'

import withFeatureRequirement from '../../lib/components/with-feature-requirement'

import sharedMessages from '../../../lib/shared-messages'
import { mayViewApplicationDevices } from '../../lib/feature-checks'
import PropTypes from '../../../lib/prop-types'

import { selectSelectedApplicationId } from '../../store/selectors/applications'

import Device from '../device'
import DeviceImport from '../device-import'
import DeviceAdd from '../device-add'
import DeviceList from '../device-list'

@connect(state => ({ appId: selectSelectedApplicationId(state) }))
@withFeatureRequirement(mayViewApplicationDevices, {
  redirect: ({ appId }) => `/applications/${appId}`,
})
@withBreadcrumb('devices', function({ appId }) {
  return <Breadcrumb path={`/applications/${appId}/devices`} content={sharedMessages.devices} />
})
export default class Devices extends React.Component {
  static propTypes = {
    match: PropTypes.match.isRequired,
  }

  render() {
    const { path } = this.props.match
    return (
      <Switch>
        <Route path={`${path}/add`} component={DeviceAdd} />
        <Route path={`${path}/import`} component={DeviceImport} />
        <Route path={`${path}/:devId`} component={Device} />
        <Route path={`${path}`} component={DeviceList} />
      </Switch>
    )
  }
}
